package skillResult

import (
	"bhakti-buana-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_IndexResult struct {
	ID        primitive.ObjectID `json:"id"`
	Code      string             `json:"code"`
	Order     int64              `json:"order"`
	Name      string             `json:"name"`
	SourceUrl string             `json:"source_url"`
	ImageUrl  string             `json:"image_url"`
	CreatedAt time.Time          `json:"created_at"`
}

// Skill Index Result
/*
 * @param skills []models.Resumes
 * @returns []S_IndexResult
 */
func Index(skills []models.Skills) []S_IndexResult {
	var results []S_IndexResult

	for _, item := range skills {
		results = append(results, S_IndexResult{
			ID:        item.ID,
			Code:      item.Code,
			Order:     item.Order,
			Name:      item.Name,
			SourceUrl: item.SourceUrl,
			ImageUrl:  item.ImageUrl,
			CreatedAt: item.CreatedAt,
		})
	}

	return results
}
