package service

import (
	"context"
	"errors"
	"strings"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/repo"
	"github.com/fatawaimamalmuftin/eventhub-be/pkg/hasing"
	jwtpkg "github.com/fatawaimamalmuftin/eventhub-be/pkg/jwt"
	"github.com/jackc/pgx/v5"
)

type LoginSrvS struct {
	Lr *repo.DbLoginRepo
}

func LoginService(lr *repo.DbLoginRepo) *LoginSrvS {
	return &LoginSrvS{
		Lr: lr,
	}
}

func (l *LoginSrvS) LoginSrv(c context.Context, logind *dto.Account) (string, error) {
	if len(logind.Password) < 6 {
		return "", cuserror.LenPas
	}

	if !strings.Contains(logind.Email, "@") {
		return "", cuserror.InvalidEmail
	}

	data, err := l.Lr.GetUserByEmail(logind.Email, c)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", cuserror.InvalidCredential
		}
		return "", cuserror.InternalError
	}

	if e := hasing.ComparePassHash(logind.Password, data.Password); e != nil {
		if errors.Is(e, cuserror.PasMissMach) {
			return "", cuserror.InvalidCredential
		}
		return "", e
	}

	// plan handler role for admin, organizer and user

	// make token
	claims := jwtpkg.NewJWTclem(data.ID, "")

	token, err := claims.GenToken()

	if err != nil {
		return "", cuserror.InternalError
	}

	return token, nil
}
