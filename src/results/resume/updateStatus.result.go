package resumeResult

import (
	"bhakti-buana-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_UpdateStatusResult struct {
	ID primitive.ObjectID `json:"id"`
}

// Resume UpdateStatus Result
/*
 * @param resume *models.Resumes
 * @returns S_UpdateStatusResult
 */
func UpdateStatus(resume *models.Resumes) S_UpdateStatusResult {
	return S_UpdateStatusResult{
		ID: resume.ID,
	}
}
