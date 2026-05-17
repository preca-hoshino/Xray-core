package trafficstats

import (
	"context"
	"net"
	"sync"

	"github.com/xtls/xray-core/common/errors"
)

// ConnTracker tracks active connections per user for kick capability.
type ConnTracker interface {
	RegisterConn(authID string, conn net.Conn)
	UnregisterConn(authID string, conn net.Conn)
	KickUser(authID string) int
	OnlineCount(authID string) int32
	OnlineUsers() map[string]int32
}

// ConnTrackerType returns the type for feature lookup.
func ConnTrackerType() interface{} {
	return (*ConnTracker)(nil)
}

// connTracker implements ConnTracker.
type connTracker struct {
	mu    sync.Mutex
	conns map[string][]net.Conn // authID → connections
}

// NewConnTracker creates a new ConnTracker.
func NewConnTracker() ConnTracker {
	return &connTracker{
		conns: make(map[string][]net.Conn),
	}
}

func (ct *connTracker) RegisterConn(authID string, conn net.Conn) {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	ct.conns[authID] = append(ct.conns[authID], conn)
}

func (ct *connTracker) UnregisterConn(authID string, conn net.Conn) {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	conns := ct.conns[authID]
	for i, c := range conns {
		if c == conn {
			ct.conns[authID] = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	if len(ct.conns[authID]) == 0 {
		delete(ct.conns, authID)
	}
}

func (ct *connTracker) KickUser(authID string) int {
	ct.mu.Lock()
	conns := ct.conns[authID]
	delete(ct.conns, authID)
	ct.mu.Unlock()

	count := 0
	for _, c := range conns {
		if err := c.Close(); err != nil {
			errors.LogInfo(context.Background(), "app/trafficstats: kick close error for ", authID, ": ", err)
		}
		count++
	}
	return count
}

func (ct *connTracker) OnlineCount(authID string) int32 {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	return int32(len(ct.conns[authID]))
}

func (ct *connTracker) OnlineUsers() map[string]int32 {
	ct.mu.Lock()
	defer ct.mu.Unlock()
	result := make(map[string]int32)
	for id, conns := range ct.conns {
		if len(conns) > 0 {
			result[id] = int32(len(conns))
		}
	}
	return result
}
