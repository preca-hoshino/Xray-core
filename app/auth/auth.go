package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/xtls/xray-core/common/errors"
)

// Authenticator is the authentication interface used by inbound handlers.
// It maps directly to the interface specified in auth-http-api-spec.md §2.
type Authenticator interface {
	Authenticate(remoteAddr net.Addr, credential string, txRate uint64) (ok bool, id string)
}

// AuthenticatorType returns the type of Authenticator for feature lookup.
func AuthenticatorType() interface{} {
	return (*Authenticator)(nil)
}

// Config is defined in config.pb.go (protobuf)

type cacheEntry struct {
	authID    string
	expiresAt time.Time
}

// HTTPAuthenticator implements Authenticator and features.Feature.
// It calls the external auth center via POST /api/v1/auth/check.
type HTTPAuthenticator struct {
	url        string
	nodeToken  string
	nodeID     string
	protocol   string
	httpClient *http.Client
	cache      sync.Map // credential(string) → cacheEntry
}

// NewAuthenticator creates a new HTTPAuthenticator from config.
func NewAuthenticator(cfg *Config) *HTTPAuthenticator {
	timeout := time.Duration(cfg.Timeout) * time.Second
	if timeout <= 0 {
		timeout = 5 * time.Second
	}
	return &HTTPAuthenticator{
		url:       cfg.AuthUrl,
		nodeToken: cfg.NodeToken,
		nodeID:    cfg.NodeId,
		protocol:  cfg.Protocol,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Type implements common.HasType.
func (a *HTTPAuthenticator) Type() interface{} {
	return AuthenticatorType()
}

// Start implements common.Runnable.
func (a *HTTPAuthenticator) Start() error {
	errors.LogInfo(context.Background(), "app/auth: HTTP authenticator started, url=", a.url, " node=", a.nodeID, " protocol=", a.protocol)
	return nil
}

// Close implements common.Closable.
func (a *HTTPAuthenticator) Close() error {
	a.httpClient.CloseIdleConnections()
	return nil
}

// Authenticate implements the Authenticator interface.
// It checks local cache first, then POSTs to the auth center.
// Non-200 status codes and errors are treated as rejection.
func (a *HTTPAuthenticator) Authenticate(remoteAddr net.Addr, credential string, txRate uint64) (bool, string) {
	// Level 1: Local cache (0 network requests on hit)
	if v, ok := a.cache.Load(credential); ok {
		e := v.(cacheEntry)
		if time.Now().Before(e.expiresAt) {
			return true, e.authID
		}
		// Entry expired — delete only if it's the same pointer (avoid
		// racing with another goroutine that may have just refreshed it).
		a.cache.CompareAndDelete(credential, v)
	}

	// Level 2: POST to auth center
	body, _ := json.Marshal(map[string]interface{}{
		"remote_addr": remoteAddr.String(),
		"credential":  credential,
		"tx":          txRate,
		"protocol":    a.protocol,
		"node_id":     a.nodeID,
		"timestamp":   time.Now().Unix(),
	})

	req, err := http.NewRequest("POST", a.url, bytes.NewReader(body))
	if err != nil {
		errors.LogInfo(context.Background(), "app/auth: failed to create request: ", err)
		return false, ""
	}
	req.Header.Set("Content-Type", "application/json")
	if a.nodeToken != "" {
		req.Header.Set("Authorization", "Bearer "+a.nodeToken)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		if err != nil {
			errors.LogInfo(context.Background(), "app/auth: auth center unreachable: ", err)
		} else {
			errors.LogInfo(context.Background(), "app/auth: auth center returned ", resp.StatusCode)
			resp.Body.Close()
		}
		return false, ""
	}
	defer resp.Body.Close()

	// Limit response size to prevent abuse
	limitedReader := io.LimitReader(resp.Body, 4096)
	var result struct {
		OK  bool   `json:"ok"`
		ID  string `json:"id"`
		TTL int64  `json:"ttl"`
	}
	if err := json.NewDecoder(limitedReader).Decode(&result); err != nil {
		errors.LogInfo(context.Background(), "app/auth: failed to decode auth response: ", err)
		return false, ""
	}

	// Level 3: Write cache on successful auth with positive TTL
	if result.OK && result.TTL > 0 {
		a.cache.Store(credential, cacheEntry{
			authID:    result.ID,
			expiresAt: time.Now().Add(time.Duration(result.TTL) * time.Second),
		})
	}

	return result.OK, result.ID
}
