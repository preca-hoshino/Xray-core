package trafficstats

import (
	"context"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/features/stats"
)

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		cfg := config.(*Config)
		if cfg.Listen == "" {
			cfg.Listen = ":9999"
		}

		srv, err := NewServer(cfg)
		if err != nil {
			return nil, errors.New("app/trafficstats: failed to create server").Base(err)
		}

		// Wire up stats.Manager (essential feature, always available at this point)
		v := core.MustFromContext(ctx)
		if sm := v.GetFeature(stats.ManagerType()); sm != nil {
			srv.SetStatsManager(sm.(stats.Manager))
		}

		return srv, nil
	}))
}
