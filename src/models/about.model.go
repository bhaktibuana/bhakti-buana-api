package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type S_AboutLocation struct {
	City     string `bson:"city" json:"city"`
	Province string `bson:"province" json:"province"`
	Country  string `bson:"country" json:"country"`
}

type S_AboutResume struct {
	ID  primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	URL string             `bson:"url" json:"url"`
}

type About struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NickName  string             `bson:"nick_name" json:"nick_name"`
	Role      string             `bson:"role" json:"role"`
	Summary   string             `bson:"summary" json:"summary"`
	Email     string             `bson:"email" json:"email"`
	Location  S_AboutLocation    `bson:"location" json:"location"`
	Resume    S_AboutResume      `bson:"resume" json:"resume"`
	Photo     string             `bson:"photo" json:"photo"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}
