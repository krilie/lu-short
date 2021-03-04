package server

import (
	"context"
	"lu-short/component/ncfg"
	"lu-short/server/httpapi"
	"time"
)

type Service struct {
	httpApi *httpapi.HttpApi
	cfg     *ncfg.NConfig
}

func NewService(httpApi *httpapi.HttpApi, cfg *ncfg.NConfig) *Service {
	return &Service{httpApi: httpApi, cfg: cfg}
}

func (s *Service) StartService(ctx context.Context) func(waitDuration time.Duration) error {
	return s.httpApi.SetRouterAndStartHttpService(ctx, s.cfg.GetHttpCfg().Port)
}
