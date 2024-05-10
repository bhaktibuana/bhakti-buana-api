package skillResult

import (
	"bhakti-buana-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_UpdateResult struct {
	ID        primitive.ObjectID `json:"id"`
	Code      string             `json:"code"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// Skill Update Result
/*
 * @param skill *models.Skills
 * @returns S_UpdateResult
 */
func Update(skill *models.Skills) S_UpdateResult {
	return S_UpdateResult{
		ID:        skill.ID,
		Code:      skill.Code,
		UpdatedAt: skill.UpdatedAt,
	}
}
