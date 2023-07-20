package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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

	config *configStruct
)

type configStruct struct {
	Token       string `json:"Token"`
	BotPrefix   string `json:"BotPrefix"`
	GuiderToken string `json:"GuiderToken"`
	GuildID     string `json:"GuildID"`
	MongoUrl    string `json:"MongoUrl"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	GuiderToken = config.GuiderToken
	GuildID = config.GuildID
	MongoUrl = config.MongoUrl

	// MongoDB connection string
	mongoURI := MongoUrl

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Get a handle to the database and collection
	Database := client.Database("test-bot")
	Collection = Database.Collection("test-bot-logs")

	return nil
}
