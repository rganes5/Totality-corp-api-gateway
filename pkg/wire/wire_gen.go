// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api/handlers"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/client"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/config"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/service"
)

// Injectors from wire.go:

func InitializeAPI(cfg *config.Config) (*api.Server, error) {
	clients, err := service.InitClient(cfg)
	if err != nil {
		return nil, err
	}
	userClient := client.NewUserClient(clients)
	userHandler := handlers.NewUserHandler(userClient)
	server, err := api.NewServerHTTP(cfg, userHandler)
	if err != nil {
		return nil, err
	}
	return server, nil
}
