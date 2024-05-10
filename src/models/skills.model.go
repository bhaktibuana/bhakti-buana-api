package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Skills struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Code      string             `bson:"code" json:"code"`
	Order     int64              `bson:"order" json:"order"`
	Name      string             `bson:"name" json:"name"`
	SourceUrl string             `bson:"source_url" json:"source_url"`
	ImageUrl  string             `bson:"image_url" json:"image_url"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}
