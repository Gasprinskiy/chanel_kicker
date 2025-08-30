package chanel_kicker

type KickUserParam struct {
	TgID   int64
	Reason int64
}

func NewKickUserParamWithSubscriptionExpireReason(tgID int64) KickUserParam {
	return KickUserParam{
		TgID:   tgID,
		Reason: KickReasonSubscriptionExpire,
	}
}
