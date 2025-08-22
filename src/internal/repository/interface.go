package repository

import "context"

type Kicker interface {
	KickExpiredSubsUsers(ctx context.Context, tgIDList []int64) error
}
