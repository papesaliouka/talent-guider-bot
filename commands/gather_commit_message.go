package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Define the channel IDs for commit messages
	commitChannelIDs := []string{
		"1131623698646958223",
		"1131619853300678798",
		"1130175812977565696",
		"1131935365469585438",
	}

	// Check if the message is from one of the commit message channels
	if containsBis(commitChannelIDs, m.ChannelID) && len(m.Message.Embeds) > 0 {
		embed := m.Message.Embeds[0]

		// Extract username, title (repoBranchLine), and commits from the embed
		username := embed.Author.Name
		description := embed.Description
		title := embed.Title

		// Call the saveCommitMessage function to save the data to MongoDB
		err := saveCommitMessage(CommitMessage{
			Username:    username,
			Title:       title,
			Description: description,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func containsBis(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

type CommitMessage struct {
	Username    string `json:"username"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func saveCommitMessage(commitMsg CommitMessage) error {
	collection := Database.Collection("commits")

	// Insert the data into MongoDB
	_, err := collection.InsertOne(context.Background(), bson.M{
		"username":    commitMsg.Username,
		"title":       commitMsg.Title,
		"description": commitMsg.Description,
	})
	if err != nil {
		log.Printf("Failed to insert data into MongoDB: %v", err)
		return err
	}

	return nil
}
