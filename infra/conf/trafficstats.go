package conf

import (
	"github.com/xtls/xray-core/app/trafficstats"
)

type TrafficStatsConfig struct {
	Listen string `json:"listen"`
	Secret string `json:"secret"`
}

// Build implements Buildable.
func (c *TrafficStatsConfig) Build() (*trafficstats.Config, error) {
	if c.Listen == "" {
		c.Listen = ":9999"
	}
	return &trafficstats.Config{
		Listen: c.Listen,
		Secret: c.Secret,
	}, nil
}
