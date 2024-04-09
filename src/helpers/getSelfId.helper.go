package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetSelfID(context *gin.Context) (string, error) {
	tokenPayload, ok := context.Get("token_payload")
	if !ok {
		return "", errors.New("token_payload not found in context")
	}

	payload, ok := tokenPayload.(jwt.MapClaims)
	if !ok {
		return "", errors.New("unable to extract ID from token payload")
	}

	id, ok := payload["id"].(string)
	if !ok {
		return "", errors.New("ID not found or not a string in token payload")
	}

	return id, nil
}
