package blacklist

import "sync"

var (
	blacklistToken = make(map[string]struct{})
	mutex          sync.RWMutex
)

func AddToken(token string) {
	mutex.Lock()
	defer mutex.Unlock()

	blacklistToken[token] = struct{}{}
}

func IsTokenBlackList(token string) bool {
	mutex.RLock()
	defer mutex.RUnlock()

	_, exist := blacklistToken[token]

	return exist
}
