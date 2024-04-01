package repository

import (
	"context"
	"fmt"
	"log"

	"Tach/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRepository struct {
	collection *mongo.Collection
}

func NewAccountRepository(client *mongo.Client, dbName, collectionName string) *AccountRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &AccountRepository{collection}
}

func (ar *AccountRepository) GetAllAccounts(ctx context.Context) ([]models.Account, error) {
	var accounts []models.Account

	cursor, err := ar.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error getting all accounts:", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var account models.Account
		if err := cursor.Decode(&account); err != nil {
			log.Println("Error decoding account:", err)
			continue
		}
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, err
	}

	return accounts, nil

}

func (ar *AccountRepository) CreateAccount(ctx context.Context, account *models.Account) error {
	_, err := ar.collection.InsertOne(ctx, account)
	if err != nil {
		log.Println("Error creating account:", err)
		return err
	}
	return nil
}

func (ar *AccountRepository) GetAccountByID(ctx context.Context, accountID string) (*models.Account, error) {
	var account models.Account
	objID, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		log.Println("Error converting ID to ObjectId:", err)
		return nil, err
	}

	filter := bson.M{"_id": objID}
	err = ar.collection.FindOne(ctx, filter).Decode(&account)
	if err != nil {
		log.Println("Error getting account by ID:", err)
		return nil, err
	}
	return &account, nil
}

func (ar *AccountRepository) UpdateAccount(accountID string, name string, balance float64, currency string) error {
	objectID, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"name":     name,
			"balance":  balance,
			"currency": currency,
		},
	}

	updateResult, err := ar.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if updateResult.ModifiedCount != 1 {
		return fmt.Errorf("could not update account")
	}

	return nil
}

func (ar *AccountRepository) DeleteAccount(accountID string) error {
	objectID, err := primitive.ObjectIDFromHex(accountID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	_, err = ar.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Error deleting account:", err)
		return err
	}
	return nil
}
