package skillResult

import (
	"bhakti-buana-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_StoreResult struct {
	ID        primitive.ObjectID `json:"id"`
	Code      string             `json:"code"`
	Order     int64              `json:"order"`
	CreatedAt time.Time          `json:"created_at"`
}

// Skill Store Result
/*
 * @param skill *models.Skills
 * @returns S_StoreResult
 */
func Store(skill *models.Skills) S_StoreResult {
	return S_StoreResult{
		ID:        skill.ID,
		Code:      skill.Code,
		Order:     skill.Order,
		CreatedAt: skill.CreatedAt,
	}
}
