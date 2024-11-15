package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jsusmachaca/tiksup/pkg/auth"
)

func CreateToken(id string, username string) (string, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":  id,
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenstring string) (jwt.MapClaims, error) {
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, auth.ErrInvalidToken
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, nil
}
