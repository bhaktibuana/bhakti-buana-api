package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	AccountType string             `bson:"account_type" json:"account_type"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt   time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}

const (
	USER_ACCOUNT_TYPE_ADMIN string = "super_admin"
	USER_ACCOUNT_TYPE_USER  string = "user"
	USER_STATUS_VERIFIED    string = "verified"
	USER_STATUS_UNVERIFIED  string = "unverified"
)
