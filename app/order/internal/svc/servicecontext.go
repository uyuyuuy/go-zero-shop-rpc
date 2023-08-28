package svc

import "github.com/uyuyuuy/go-zero-shop-rpc/app/order/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
