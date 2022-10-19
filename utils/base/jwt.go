package base

import (
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken 生成Token
func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(key))
}
