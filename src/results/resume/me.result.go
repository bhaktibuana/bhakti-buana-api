package resumeResult

import (
	"bhakti-buana-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_StoreResult struct {
	ID        primitive.ObjectID `json:"id"`
	Title     string             `json:"title"`
	URL       string             `json:"url"`
	Status    string             `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
}

// Store Request
/*
 * @param user *models.Resumes
 * @returns S_StoreResult
 */
func Store(resume *models.Resumes) S_StoreResult {
	return S_StoreResult{
		ID:        resume.ID,
		Title:     resume.Title,
		URL:       resume.URL,
		Status:    resume.Status,
		CreatedAt: resume.CreatedAt,
	}
}
