package authResult

import (
	"bhakti-buana-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_MeResult struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	AccountType string             `json:"account_type"`
	Status      string             `json:"status"`
}

// Me Request
/*
 * @param user *models.Users
 * @returns S_MeResult
 */
func Me(user *models.Users) S_MeResult {
	return S_MeResult{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AccountType: user.AccountType,
		Status:      user.Status,
	}
}
