package trafficstats

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/features"
	"github.com/xtls/xray-core/features/stats"
)

// Server is the traffic statistics HTTP API server.
// It implements features.Feature and ConnTracker.
type Server struct {
	listen  string
	secret  string
	server  *http.Server
	stats   stats.Manager // injected via SetStatsManager before Start()
	tracker *connTracker

	started atomic.Bool
}

// NewServer creates a new traffic stats server.
func NewServer(cfg *Config) (*Server, error) {
	s := &Server{
		listen:  cfg.Listen,
		secret:  cfg.Secret,
		tracker: &connTracker{conns: make(map[string][]net.Conn)},
	}
	return s, nil
}

// Type implements common.HasType.
func (s *Server) Type() interface{} {
	// Return both ConnTrackerType and Feature
	return ConnTrackerType()
}

// Start implements common.Runnable.
func (s *Server) Start() error {
	if s.started.Swap(true) {
		return nil
	}

	// We need stats.Manager from the core instance. Since we can't access the core
	// instance directly in Start(), we rely on it being injected beforehand.
	// If stats is nil, we operate without traffic stats (only online/kick work).
	// The stats manager will be injected via SetStatsManager before Start().

	mux := http.NewServeMux()
	mux.HandleFunc("/traffic", s.authMiddleware(s.handleTraffic))
	mux.HandleFunc("/online", s.authMiddleware(s.handleOnline))
	mux.HandleFunc("/kick", s.authMiddleware(s.handleKick))

	s.server = &http.Server{
		Addr:    s.listen,
		Handler: mux,
	}

	ln, err := net.Listen("tcp", s.listen)
	if err != nil {
		return errors.New("app/trafficstats: failed to listen on ", s.listen).Base(err)
	}

	go func() {
		if err := s.server.Serve(ln); err != nil && err != http.ErrServerClosed {
			errors.LogError(context.Background(), "app/trafficstats: HTTP server error: ", err)
		}
	}()

	errors.LogInfo(context.Background(), "app/trafficstats: HTTP server started on ", s.listen)
	return nil
}

// Close implements common.Closable.
func (s *Server) Close() error {
	if s.server != nil {
		return s.server.Close()
	}
	return nil
}

// SetStatsManager injects the stats manager. Called by config creator.
func (s *Server) SetStatsManager(sm stats.Manager) {
	s.stats = sm
}

// ConnTracker delegation

func (s *Server) RegisterConn(authID string, conn net.Conn) {
	s.tracker.RegisterConn(authID, conn)
}

func (s *Server) UnregisterConn(authID string, conn net.Conn) {
	s.tracker.UnregisterConn(authID, conn)
}

func (s *Server) KickUser(authID string) int {
	return s.tracker.KickUser(authID)
}

func (s *Server) OnlineCount(authID string) int32 {
	return s.tracker.OnlineCount(authID)
}

func (s *Server) OnlineUsers() map[string]int32 {
	return s.tracker.OnlineUsers()
}

// Middleware

func (s *Server) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.secret == "" {
			next(w, r)
			return
		}
		auth := r.Header.Get("Authorization")
		if subtle.ConstantTimeCompare([]byte(auth), []byte(s.secret)) != 1 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next(w, r)
	}
}

// GET /traffic and GET /traffic?clear=1
func (s *Server) handleTraffic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	clear := r.URL.Query().Get("clear") == "1"
	result := make(map[string]map[string]uint64)

	if s.stats != nil {
		s.stats.VisitCounters(func(name string, counter stats.Counter) bool {
			authID := extractAuthID(name)
			if authID == "" {
				return true
			}
			val := uint64(counter.Value())
			if entry, ok := result[authID]; ok {
				if strings.HasSuffix(name, ">>>uplink") {
					entry["tx"] += val
				} else if strings.HasSuffix(name, ">>>downlink") {
					entry["rx"] += val
				}
			} else {
				entry = map[string]uint64{"tx": 0, "rx": 0}
				if strings.HasSuffix(name, ">>>uplink") {
					entry["tx"] = val
				} else if strings.HasSuffix(name, ">>>downlink") {
					entry["rx"] = val
				}
				result[authID] = entry
			}

			if clear {
				counter.Set(0)
			}
			return true
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GET /online
func (s *Server) handleOnline(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	result := s.tracker.OnlineUsers()

	// Also merge in online maps from stats manager
	if s.stats != nil {
		s.stats.VisitOnlineMaps(func(name string, om stats.OnlineMap) bool {
			authID := extractAuthID(name)
			if authID == "" {
				return true
			}
			count := om.Count()
			if count > 0 {
				existing := result[authID]
				if int(existing) < count {
					result[authID] = int32(count)
				}
			}
			return true
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// POST /kick
func (s *Server) handleKick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Limit request body to 64KB to prevent abuse
	limitedBody := io.LimitReader(r.Body, 64*1024)
	var ids []string
	if err := json.NewDecoder(limitedBody).Decode(&ids); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if len(ids) > 1000 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "too many IDs (max 1000)"})
		return
	}

	kicked := 0
	for _, id := range ids {
		kicked += s.tracker.KickUser(id)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"kicked": kicked})
}

// extractAuthID extracts the user email/authID from a stats counter/online-map name.
// Format: "user>>>EMAIL>>>traffic>>>uplink" or "user>>>EMAIL>>>online"
func extractAuthID(name string) string {
	const prefix = "user>>>"
	if !strings.HasPrefix(name, prefix) {
		return ""
	}
	rest := name[len(prefix):]
	idx := strings.Index(rest, ">>>")
	if idx < 0 {
		return ""
	}
	return rest[:idx]
}

// Ensure Server implements ConnTracker
var _ ConnTracker = (*Server)(nil)

// Ensure Server implements features.Feature
var _ features.Feature = (*Server)(nil)
