package service

import (
	"strings"

	"github.com/fatawaimamalmuftin/eventhub-be/pkg/blacklist"
)

type LogoutSrvS struct {
}

func LogoutService() *LogoutSrvS {
	return &LogoutSrvS{}
}

func (l *LogoutSrvS) Logout(authorization string) error {
	bearer := strings.Split(authorization, " ")

	blacklist.AddToken(bearer[1])

	return nil
}
