package grpc_client

import (
	"chanel_kicker/src/internal/repository"
	"chanel_kicker/src/internal/repository/grpc_client/proto/kicker"
	"context"

	"google.golang.org/grpc"
)

type kickerRepo struct {
	client kicker.KickerServiceClient
}

func NewKicker(conn *grpc.ClientConn) repository.Kicker {
	return &kickerRepo{
		kicker.NewKickerServiceClient(conn),
	}
}

func (r *kickerRepo) KickExpiredSubsUsers(tgIDList []int64) error {
	_, err := r.client.KickExpiredSubsUsers(context.Background(), &kicker.KickExpiredSubsUsersRequest{
		TgIds: tgIDList,
	})

	return err
}
