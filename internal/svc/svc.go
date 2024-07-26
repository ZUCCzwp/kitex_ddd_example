package svc

import "github.com/ZUCCzwp/ddd/my-awesome-service/internal/config"

type ServiceContext struct {
	config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		config: c,
	}
}
