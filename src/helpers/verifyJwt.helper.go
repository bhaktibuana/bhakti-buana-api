package helpers

import (
	"bhakti-buana-api/src/configs"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func VerifyJwt(tokenString string) (jwt.MapClaims, error) {
	var err error

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.AppConfig().JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("failed to parse claims")
		}
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, fmt.Errorf("token is malformed")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, fmt.Errorf("token is expired or not valid yet")
		} else {
			return nil, fmt.Errorf("token validation failed")
		}
	} else {
		return nil, fmt.Errorf("token is not valid")
	}
}
