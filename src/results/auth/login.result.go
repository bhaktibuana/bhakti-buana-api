package authResult

import (
	"bhakti-buana-api/src/helpers"
	"bhakti-buana-api/src/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_LoginResult struct {
	ID    primitive.ObjectID `json:"id"`
	Token string             `json:"token"`
}

// Auth Login Result
/*
 * @param user *models.Users
 * @returns S_LoginResult
 */
func Login(user *models.Users) S_LoginResult {

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	}

	token, _ := helpers.GenerateJWT(claims, time.Hour*24*30)

	return S_LoginResult{
		ID:    user.ID,
		Token: token,
	}
}
