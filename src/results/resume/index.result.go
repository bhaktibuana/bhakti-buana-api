package resumeResult

import (
	"bhakti-buana-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_IndexResult struct {
	ID         primitive.ObjectID `json:"id"`
	Title      string             `json:"title"`
	URL        string             `json:"url"`
	Downloaded int64              `json:"downloaded"`
	Status     string             `json:"status"`
	CreatedAt  time.Time          `json:"created_at"`
}

// Resume Index Result
/*
 * @param resumes []models.Resumes
 * @returns []S_IndexResult
 */
func Index(resumes []models.Resumes) []S_IndexResult {
	var results []S_IndexResult

	for _, item := range resumes {
		results = append(results, S_IndexResult{
			ID:         item.ID,
			Title:      item.Title,
			URL:        item.URL,
			Downloaded: item.Downloaded,
			Status:     item.Status,
			CreatedAt:  item.CreatedAt,
		})
	}

	return results
}
