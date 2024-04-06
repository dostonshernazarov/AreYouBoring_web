package services

import (

	boredapi "github/NotBoring/bored_api"
	"github/NotBoring/config"

	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	BoredService() boredapi.BoredServiceClient
}

type serviceManager struct {
	boredService boredapi.BoredServiceClient
}

func (s *serviceManager) BoredService() boredapi.BoredServiceClient {
	return s.boredService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	

	serviceManager := &serviceManager{
		boredService: boredapi.NewBoredClient(),
	}

	return serviceManager, nil
}
