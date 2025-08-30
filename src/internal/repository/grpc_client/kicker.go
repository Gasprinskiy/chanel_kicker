package grpc_client

import (
	"chanel_kicker/src/internal/entity/chanel_kicker"
	"chanel_kicker/src/internal/repository"
	"chanel_kicker/src/internal/repository/grpc_client/proto/kicker"
	"chanel_kicker/src/tools/slice"
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

func (r *kickerRepo) KickExpiredSubsUsers(params []chanel_kicker.KickUserParam) error {
	_, err := r.client.KickUsers(context.Background(), &kicker.KickUsersRequest{
		Params: slice.Map(params, func(item chanel_kicker.KickUserParam) *kicker.KickUserParam {
			return &kicker.KickUserParam{
				TgId:     item.TgID,
				ReasonId: item.Reason,
			}
		}),
	})

	return err
}
