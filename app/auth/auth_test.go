package auth

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProtoRoundTrip(t *testing.T) {
	orig := &Config{
		AuthUrl:   "https://example.com/auth",
		NodeToken: "mytoken",
		NodeId:    "node-01",
		Protocol:  "vless",
		Timeout:   5,
	}

	data, err := proto.Marshal(orig)
	if err != nil {
		t.Fatalf("proto.Marshal failed: %v", err)
	}

	decoded := &Config{}
	if err := proto.Unmarshal(data, decoded); err != nil {
		t.Fatalf("proto.Unmarshal failed: %v", err)
	}

	if decoded.AuthUrl != orig.AuthUrl {
		t.Errorf("AuthUrl mismatch: got %q, want %q", decoded.AuthUrl, orig.AuthUrl)
	}
	if decoded.NodeToken != orig.NodeToken {
		t.Errorf("NodeToken mismatch: got %q, want %q", decoded.NodeToken, orig.NodeToken)
	}
	if decoded.NodeId != orig.NodeId {
		t.Errorf("NodeId mismatch: got %q, want %q", decoded.NodeId, orig.NodeId)
	}
	if decoded.Protocol != orig.Protocol {
		t.Errorf("Protocol mismatch: got %q, want %q", decoded.Protocol, orig.Protocol)
	}
	if decoded.Timeout != orig.Timeout {
		t.Errorf("Timeout mismatch: got %d, want %d", decoded.Timeout, orig.Timeout)
	}
}
