package middlewares

import (
	"bhakti-buana-api/src/helpers"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuthenticate(context *gin.Context) {
	authorizationHeader := context.GetHeader("Authorization")

	if authorizationHeader == "" {
		helpers.HttpResponse("Unauthorized", http.StatusUnauthorized, context, nil)
		return
	}

	splitToken := strings.Split(authorizationHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		helpers.HttpResponse("Wrong authorization format", http.StatusBadRequest, context, nil)
		return
	}

	tokenString := splitToken[1]
	claims, err := helpers.VerifyJwt(tokenString)
	if err != nil {
		if validationError, ok := err.(*jwt.ValidationError); ok && validationError.Errors&jwt.ValidationErrorExpired != 0 {
			helpers.HttpResponse("Unauthorized: Token expired", http.StatusUnauthorized, context, nil)
			return
		}
		helpers.HttpResponse("Unauthorized: Invalid token", http.StatusUnauthorized, context, nil)
		return
	}

	context.Set("token_payload", claims)
	context.Set("token_string", tokenString)
	context.Next()
}
