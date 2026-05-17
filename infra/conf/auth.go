package conf

import (
	"github.com/xtls/xray-core/app/auth"
)

type AuthConfig struct {
	AuthUrl   string `json:"authUrl"`
	NodeToken string `json:"nodeToken"`
	NodeId    string `json:"nodeId"`
	Protocol  string `json:"protocol"`
	Timeout   int64  `json:"timeout"`
}

// Build implements Buildable.
func (c *AuthConfig) Build() (*auth.Config, error) {
	if c.Protocol == "" {
		c.Protocol = "vless"
	}
	return &auth.Config{
		AuthUrl:   c.AuthUrl,
		NodeToken: c.NodeToken,
		NodeId:    c.NodeId,
		Protocol:  c.Protocol,
		Timeout:   c.Timeout,
	}, nil
}
