package repository

import (
	"chanel_kicker/src/internal/transaction"
)

type Kicker interface {
	KickExpiredSubsUsers(tgIDList []int64) error
}

type Purchases interface {
	LoadExpiredPurchasesUserIDS(ts transaction.Session) ([]int64, error)
}
