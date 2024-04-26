package aboutResult

import (
	"bhakti-buana-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_UpdateResult struct {
	ID primitive.ObjectID `json:"id"`
}

// About Update Result
/*
 * @param resume *models.About
 * @returns S_UpdateResult
 */
func Update(resume *models.About) S_UpdateResult {
	return S_UpdateResult{
		ID: resume.ID,
	}
}
