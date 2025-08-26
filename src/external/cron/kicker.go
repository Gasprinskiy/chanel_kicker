package cron

import (
	"chanel_kicker/src/internal/transaction"
	"chanel_kicker/src/tools/logger"
	"chanel_kicker/src/uimport"
)

type KickerCron struct {
	log *logger.Logger
	sm  transaction.SessionManager
	ui  *uimport.UsecaseImport
}

func NewKickerCron(
	log *logger.Logger,
	sm transaction.SessionManager,
	ui *uimport.UsecaseImport,
) *KickerCron {
	cron := KickerCron{
		log,
		sm,
		ui,
	}

	return &cron
}

func (c *KickerCron) KickExpiredSubsUsers() {
	ts := c.sm.CreateSession()
	if err := ts.Start(); err != nil {
		c.log.Db.Errorln("не удалось запустить транзакцию")
		return
	}

	defer ts.Rollback()

	c.ui.Kicker.KickExpiredSubsUsers(ts)
}
