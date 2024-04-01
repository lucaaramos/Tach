package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func SetupDB() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		return fmt.Errorf("MONGO_URL environment variable not defined")
	}

	clientOptions := options.Client().ApplyURI(mongoURL)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("Error conecting to the data base: %v", err)
	}
	fmt.Println("Connection to database established successfully")

	return nil
}

func GetClient() *mongo.Client {
	return client
}
