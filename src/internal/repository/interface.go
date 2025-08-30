package repository

import (
	"chanel_kicker/src/internal/entity/chanel_kicker"
	"chanel_kicker/src/internal/transaction"
)

type Kicker interface {
	KickExpiredSubsUsers(params []chanel_kicker.KickUserParam) error
}

type Purchases interface {
	LoadExpiredPurchasesUserIDS(ts transaction.Session) ([]int64, error)
}
