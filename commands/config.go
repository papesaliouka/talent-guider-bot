package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Token       string
	BotPrefix   string
	GuiderToken string
	GuildID     string
	MongoUrl    string
	Collection  *mongo.Collection
	Database    *mongo.Database
)



func ReadConfig() error {
	fmt.Println("Reading config from environment variables...")

	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read the environment variables
	Token = os.Getenv("Token")
	BotPrefix = os.Getenv("BotPrefix")
	GuiderToken = os.Getenv("GuiderToken")
	GuildID = os.Getenv("GuildID")
	MongoUrl = os.Getenv("MongoUrl")

	// MongoDB connection string
	mongoURI := MongoUrl

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Get a handle to the database and collection
	Database = client.Database("test-bot")
	Collection = Database.Collection("test-bot-logs")

	return nil
}
