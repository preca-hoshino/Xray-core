package trafficstats

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProtoRoundTrip(t *testing.T) {
	orig := &Config{
		Listen: "127.0.0.1:9999",
		Secret: "mysecret",
	}

	data, err := proto.Marshal(orig)
	if err != nil {
		t.Fatalf("proto.Marshal failed: %v", err)
	}
	t.Logf("Marshaled %d bytes", len(data))

	decoded := &Config{}
	if err := proto.Unmarshal(data, decoded); err != nil {
		t.Fatalf("proto.Unmarshal failed: %v", err)
	}

	if decoded.Listen != orig.Listen {
		t.Errorf("Listen mismatch: got %q, want %q", decoded.Listen, orig.Listen)
	}
	if decoded.Secret != orig.Secret {
		t.Errorf("Secret mismatch: got %q, want %q", decoded.Secret, orig.Secret)
	}
}
