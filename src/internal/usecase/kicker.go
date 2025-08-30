package usecase

import (
	"chanel_kicker/src/internal/entity/chanel_kicker"
	"chanel_kicker/src/internal/entity/global"
	"chanel_kicker/src/internal/transaction"
	"chanel_kicker/src/rimport"
	"chanel_kicker/src/tools/logger"
	"chanel_kicker/src/tools/slice"
)

type Kicker struct {
	log *logger.Logger
	ri  *rimport.RepositoryImports
}

func NewKicker(
	log *logger.Logger,
	ri *rimport.RepositoryImports,
) *Kicker {
	return &Kicker{log, ri}
}

func (u *Kicker) KickExpiredSubsUsers(ts transaction.Session) error {
	idList, err := u.ri.Repository.Purchases.LoadExpiredPurchasesUserIDS(ts)
	switch err {
	case nil:
	case global.ErrNoData:
		u.log.File.Info("нет пользователей с истекшими подписками")
		return nil
	default:
		u.log.Db.Errorln("не удалось найти пользователей с истекшими подписками, ошибка:", err)
		return global.ErrInternalError
	}

	params := slice.Map(idList, func(tgID int64) chanel_kicker.KickUserParam {
		return chanel_kicker.NewKickUserParamWithSubscriptionExpireReason(tgID)
	})
	err = u.ri.Repository.Kicker.KickExpiredSubsUsers(params)
	if err != nil {
		u.log.Db.Errorln("не удалось кикнуть пользователей с истекшими подписками, ошибка:", err)
		return global.ErrInternalError
	}

	u.log.File.Infof("кинуто %d пользователей", len(idList))
	return nil
}
