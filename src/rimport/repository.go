package rimport

import "chanel_kicker/src/internal/repository"

type Repository struct {
	Kicker    repository.Kicker
	Purchases repository.Purchases
}
