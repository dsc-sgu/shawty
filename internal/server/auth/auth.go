package auth

import (
	"github.com/dsc-sgu/shawty/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtParser *jwt.Parser = jwt.NewParser(
	jwt.WithValidMethods([]string{"HS256"}),
)

func CheckSession(session string) error {
	_, err := jwtParser.Parse(session, func(t *jwt.Token) (any, error) {
		return []byte(config.C.JwtSecret), nil
	})
	return err
}
