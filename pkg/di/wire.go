package di

import (
	"api/pkg/api"
	"api/pkg/api/handlers"
	"api/pkg/client"
	"api/pkg/config"

	"github.com/google/wire"
)

func InitializeAp(c *config.Config) (*api.Server,error) {
	wire.Build(auth.InitClient,auth.NewAuthServiceClient,handlers.NewAuthHandler,api.NewserverHttp)
	return &api.Server{},nil
}