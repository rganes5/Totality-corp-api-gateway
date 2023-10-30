package service

import (
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/config"
	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Usercli pb.UserServiceClient
}

func InitClient(cfg *config.Config) (Clients, error) {
	usercc, userErr := grpc.Dial(cfg.UserSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if userErr != nil {
		return Clients{}, userErr
	}

	userClient := pb.NewUserServiceClient(usercc)
	return Clients{
		Usercli: userClient,
	}, nil
}
