package service

import "github.com/fatawaimamalmuftin/eventhub-be/internal/repo"

type LoginSrvS struct {
	Lr *repo.DbLoginRepo
}

func LoginService(lr *repo.DbLoginRepo) *LoginSrvS {
	return &LoginSrvS{
		Lr: lr,
	}
}
