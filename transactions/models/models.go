package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FromAccountID   primitive.ObjectID `json:"from_account_id" bson:"from_account_id"`
	ToAccountID     primitive.ObjectID `json:"to_account_id" bson:"to_account_id"`
	Amount          float64            `json:"amount"`
	Currency        string             `json:"currency"`
	TransactionType string             `json:"transaction_type"`
	Description     string             `json:"description"`
	CreatedAt       time.Time          `json:"created_at"`
}
