package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Expense struct {
	Id          primitive.ObjectID `bson:"_id"`
	Description string             `json:"description"`
	Amount      float64            `json:"amount"`
	DueDate     time.Time          `json:"due_date"`
	IsPaid      bool               `json:"is_paid"`
	EnteredBy   *User              `json:"entered_by"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}
