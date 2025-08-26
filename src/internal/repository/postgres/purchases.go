package postgres

import (
	"chanel_kicker/src/internal/repository"
	"chanel_kicker/src/internal/transaction"
	"chanel_kicker/src/tools/sql_gen"
)

type purchasesRepo struct{}

func NewPurchases() repository.Purchases {
	return &purchasesRepo{}
}

func (r *purchasesRepo) LoadExpiredPurchasesUserIDS(ts transaction.Session) ([]int64, error) {
	sqlQuery := `
		SELECT u.tg_id
		FROM bot_users_purchases p
			JOIN bot_users_profile u ON (u.u_id = p.u_id)
			JOIN bot_subscription_types st ON (st.sub_id = p.sub_id)
		WHERE p.kick_time IS NULL
		AND DATE_TRUNC('day', CURRENT_TIMESTAMP) >= DATE_TRUNC('day', p.p_time + make_interval(months := st.term_in_month))
	`

	return sql_gen.Select[int64](SqlxTx(ts), sqlQuery)
}
