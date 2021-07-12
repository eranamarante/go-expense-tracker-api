package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	Id          primitive.ObjectID `bson:"_id"`
	Description string             `json:"description" bson:"description"`
	Amount      float64            `json:"amount" bson:"amount"`
	DueDate     time.Time          `json:"due_date" bson:"due_date"`
	IsPaid      bool               `json:"is_paid" bson:"is_paid"`
	EnteredBy   *User              `json:"entered_by" bson:"entered_by"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
