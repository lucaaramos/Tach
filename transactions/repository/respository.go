package repository

import (
	"context"
	"fmt"
	"log"

	"transactions/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepository struct {
	collection *mongo.Collection
}

func NewTransactionRepository(client *mongo.Client, dbName, collectionName string) *TransactionRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &TransactionRepository{collection}
}

func (tr *TransactionRepository) GetAllTransactions(ctx context.Context) ([]models.Transaction, error) {
	var transactions []models.Transaction

	cursor, err := tr.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error getting all transactions:", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var transaction models.Transaction
		if err := cursor.Decode(&transaction); err != nil {
			log.Println("Error decoding transaction:", err)
			continue
		}
		transactions = append(transactions, transaction)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return transactions, nil
}

func (tr *TransactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	_, err := tr.collection.InsertOne(ctx, transaction)
	if err != nil {
		log.Println("Error creating transaction:", err)
		return err
	}
	return nil
}

func (tr *TransactionRepository) GetTransactionByID(ctx context.Context, transactionID string) (*models.Transaction, error) {
	var transaction models.Transaction
	objID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		log.Println("Error converting ID to ObjectId:", err)
		return nil, err
	}

	filter := bson.M{"_id": objID}
	err = tr.collection.FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		log.Println("Error getting transaction by ID:", err)
		return nil, err
	}
	return &transaction, nil
}

func (tr *TransactionRepository) UpdateTransaction(transactionID string, updatedTransaction *models.Transaction) error {
	objectID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"from_account_id":  updatedTransaction.FromAccountID,
			"to_account_id":    updatedTransaction.ToAccountID,
			"amount":           updatedTransaction.Amount,
			"currency":         updatedTransaction.Currency,
			"transaction_type": updatedTransaction.TransactionType,
			"description":      updatedTransaction.Description,
			"created_at":       updatedTransaction.CreatedAt,
		},
	}

	updateResult, err := tr.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if updateResult.ModifiedCount != 1 {
		return fmt.Errorf("Transaction could not be updated")
	}

	return nil
}

func (tr *TransactionRepository) DeleteTransaction(ctx context.Context, transactionID string) error {
	objID, err := primitive.ObjectIDFromHex(transactionID)
	if err != nil {
		log.Println("Error converting transaction ID to ObjectID:", err)
		return err
	}
	filter := bson.M{"_id": objID}
	_, err = tr.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Error deleting transaction:", err)
		return err
	}

	fmt.Println("Transaction deleted successfully")
	return nil
}
