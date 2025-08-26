package rimport

import (
	"chanel_kicker/src/internal/repository/grpc_client"
	"chanel_kicker/src/internal/repository/postgres"

	"google.golang.org/grpc"
)

type RepositoryImports struct {
	Repository
}

func NewRepositoryImports(grpcConn *grpc.ClientConn) *RepositoryImports {
	return &RepositoryImports{
		Repository: Repository{
			Kicker:    grpc_client.NewKicker(grpcConn),
			Purchases: postgres.NewPurchases(),
		},
	}
}
