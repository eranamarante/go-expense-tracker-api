package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" validate:"required" bson:"email"`
	Password  string             `json:"password" validate:"required" bson:"password"`
	Username  string             `json:"username" bson:"username"`
	TokenHash string             `json:"token_hash" bson:"token_hash"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
