package jwt

import (
	"os"
	"time"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/golang-jwt/jwt/v5"
)

type JWTclem struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTclem(id int, role string) *JWTclem {
	return &JWTclem{
		Id:        id,
		Role:      role,
		Issuer:    os.Getenv("JWT_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
	}
}

func (j *JWTclem) GenToken() (string, error) {
	key := os.Getenv("JWT_KEY")

	if key == "" {
		return "", cuserror.ErrMissingKey
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)

	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
