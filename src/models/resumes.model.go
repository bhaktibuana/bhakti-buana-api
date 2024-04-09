package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resumes struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string             `bson:"title" json:"title"`
	URL        string             `bson:"url" json:"url"`
	Downloaded int64              `bson:"downloaded" json:"downloaded"`
	Status     string             `bson:"status" json:"status"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt  time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}

const (
	RESUME_STATUS_ACTIVE   string = "active"
	RESUME_STATUS_INACTIVE string = "inactive"
	RESUME_STATUS_DELETED  string = "deleted"
)
