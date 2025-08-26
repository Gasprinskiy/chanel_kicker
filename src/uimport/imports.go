package uimport

import (
	"chanel_kicker/src/internal/usecase"
	"chanel_kicker/src/rimport"
	"chanel_kicker/src/tools/logger"
)

type UsecaseImport struct {
	Usecase
}

func NewUsecaseImport(
	ri *rimport.RepositoryImports,
	log *logger.Logger,
) *UsecaseImport {
	return &UsecaseImport{
		Usecase: Usecase{
			Kicker: usecase.NewKicker(log, ri),
		},
	}
}
