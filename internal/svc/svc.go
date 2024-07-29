package svc

import "github.com/ZUCCzwp/kitex_ddd_example/internal/config"

type ServiceContext struct {
	config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		config: c,
	}
}
