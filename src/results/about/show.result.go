package aboutResult

import (
	"bhakti-buana-api/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_AboutLocation struct {
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}

type S_AboutResume struct {
	ID  primitive.ObjectID `json:"id"`
	URL string             `json:"url"`
}

type S_ShowResult struct {
	ID        primitive.ObjectID `json:"id"`
	NickName  string             `json:"nick_name"`
	Role      string             `json:"role"`
	Email     string             `json:"email"`
	Summary   string             `json:"summary"`
	Location  S_AboutLocation    `json:"location"`
	Resume    S_AboutResume      `json:"resume"`
	Photo     string             `json:"photo"`
	CreatedAt time.Time          `json:"created_at"`
}

// About Show Result
/*
 * @param about *models.About
 * @returns S_ShowResult
 */
func Show(about *models.About) S_ShowResult {
	return S_ShowResult{
		ID:        about.ID,
		NickName:  about.NickName,
		Role:      about.Role,
		Email:     about.Email,
		Summary:   about.Summary,
		Location:  S_AboutLocation(about.Location),
		Resume:    S_AboutResume(about.Resume),
		Photo:     about.Photo,
		CreatedAt: about.CreatedAt,
	}
}
