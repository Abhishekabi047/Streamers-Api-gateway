package di

import (
	"api/pkg/api"
	"api/pkg/api/handlers"
	auth "api/pkg/client"
	"api/pkg/config"

	"github.com/google/wire"
)

func InitializeAp(c *config.Config) (*api.Server, error) {
	wire.Build(auth.InitClient, auth.InitVideoClient, auth.NewVideoClient, auth.NewAuthServiceClient, handlers.NewAuthHandler, handlers.NewVideoHandler, api.NewserverHttp)
	return &api.Server{}, nil
}
