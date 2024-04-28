package aboutResult

import (
	"bhakti-buana-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_UpdatePhotoResult struct {
	ID primitive.ObjectID `json:"id"`
}

// About UpdUpdatePhotoate Result
/*
 * @param about *models.About
 * @returns S_UpdateResult
 */
func UpdatePhoto(about *models.About) S_UpdatePhotoResult {
	return S_UpdatePhotoResult{
		ID: about.ID,
	}
}
