package helpers

import (
	"bhakti-buana-api/src/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT Helper
/*
 * @param payload jwt.Claims
 * @param expiresIn time.Duration
 * @returns (string, error)
 */
func GenerateJWT(payload jwt.Claims, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	if expiresIn > 0 {
		token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(expiresIn).Unix()
	}
	signedToken, _ := token.SignedString([]byte(configs.AppConfig().JWT_SECRET_KEY))
	return signedToken, nil
}
