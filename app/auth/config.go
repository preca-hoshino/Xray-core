package auth

import (
	"context"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/errors"
)

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		cfg := config.(*Config)
		if cfg.AuthUrl == "" {
			return nil, errors.New("app/auth: auth_url is required")
		}
		if cfg.NodeId == "" {
			return nil, errors.New("app/auth: node_id is required")
		}
		if cfg.Protocol == "" {
			return nil, errors.New("app/auth: protocol is required")
		}
		return NewAuthenticator(cfg), nil
	}))
}
