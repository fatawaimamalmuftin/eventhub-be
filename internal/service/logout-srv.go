package service

import (
	"github.com/fatawaimamalmuftin/eventhub-be/pkg/blacklist"
)

type LogoutSrvS struct {
}

func LogoutService() *LogoutSrvS {
	return &LogoutSrvS{}
}

func (l *LogoutSrvS) Logout(token string) error {
	blacklist.AddToken(token)
	return nil
}
