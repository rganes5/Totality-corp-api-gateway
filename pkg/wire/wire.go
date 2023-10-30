//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/api/handlers"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/client"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/config"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/service"
)

func InitializeAPI(cfg *config.Config) (*api.Server, error) {
	wire.Build(service.InitClient,
		client.NewUserClient,
		handlers.NewUserHandler,
		api.NewServerHTTP)
	return &api.Server{}, nil
}
