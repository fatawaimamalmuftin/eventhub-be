package service

import (
	"context"
	"strings"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
)

type RegisSrvS struct {
	Rs *repo.DbRegisRepo
}

func RegisService(rp *repo.DbRegisRepo) *RegisSrvS {
	return &RegisSrvS{
		Rs: rp,
	}
}

func (r *RegisSrvS) RegisSrv(newUser *dto.Regis, c context.Context) error {
	if len(newUser.Password) < 6 {
		return cuserror.LenPas
	}

	if !strings.Contains(newUser.Email, "@") {
		return cuserror.InvalidEmail
	}

	if e := r.Rs.NewAccount(c, *newUser); e != nil {
		return e
	}

	return nil
}
