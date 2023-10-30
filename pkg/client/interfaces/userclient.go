package interfaces

import (
	"context"

	"github.com/rganes5/Totality-Corp/Totality-corp-api-gateway/pkg/pb"
)

type UserClient interface {
	GetUserData(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error)
	GetUserDataList(ctx context.Context, request *pb.GetUserDataListRequest) (*pb.GetUserDataListResponse, error)
}
