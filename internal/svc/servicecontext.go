package svc

import (
	"github.com/MBR2022/gosimpler/internal/config"
	"github.com/MBR2022/gosimpler/internal/middleware"
	"github.com/MBR2022/gosimpler/internal/store"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	BasicAuth rest.Middleware
	MemStore  store.MemStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		BasicAuth: middleware.NewBasicAuthMiddleware(c.BasicAuthUsername, c.BasicAuthPassword).Handle,
		MemStore:  store.NewMemStore(),
	}
}
