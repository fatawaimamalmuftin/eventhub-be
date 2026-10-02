package blacklist

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func AddToken(token string, rdb *redis.Client, c context.Context) error {
	quewe := "eventhub:token" + token
	err := rdb.Set(c, quewe, "blacklisted", 0).Err()
	return err
}

func IsTokenBlackList(token string, rdb *redis.Client, c context.Context) (bool, error) {
	quewe := "eventhub:token" + token

	err := rdb.Get(c, quewe).Err()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
