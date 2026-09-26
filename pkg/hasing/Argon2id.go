package hasing

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Memory uint32
	Time   uint32
	Thread uint8
	Key    uint32
	Salt   uint32
}

func GenerateHashConfig(
	memory,
	time,
	key,
	salt uint32,
	thread uint8,
) *HashConfig {
	return &HashConfig{
		Memory: memory,
		Time:   time,
		Thread: thread,
		Key:    key,
		Salt:   salt,
	}
}

func GenRecomConfHash() *HashConfig {
	return &HashConfig{
		Memory: 64 * 1024,
		Time:   2,
		Thread: 2,
		Key:    32,
		Salt:   16,
	}
}

func generateSalt(saltLen uint32) ([]byte, error) {
	salt := make([]byte, saltLen)

	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func (h *HashConfig) GenPasHash(pass string) (string, error) {
	salt, err := generateSalt(h.Salt)

	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(pass), salt, h.Time, h.Memory, h.Thread, h.Key,
	)

	base64hash := base64.RawStdEncoding.EncodeToString(hash)
	base64salt := base64.RawStdEncoding.EncodeToString(salt)

	completeHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, h.Memory, h.Time, h.Thread, base64salt, base64hash,
	)

	return completeHash, nil
}
