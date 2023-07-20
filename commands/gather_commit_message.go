package commands

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if the message is from the channel where commit messages are sent
	if m.ChannelID == "1131623698646958223" || m.ChannelID == "1131619853300678798" || m.ChannelID == "1130175812977565696" {
		err := saveCommitMessage(m.Content)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

type CommitMessage struct {
	Username string   `json:"username"`
	Repo     string   `json:"repository"`
	Branch   string   `json:"branch"`
	Commits  []string `json:"commits"`
}

func saveCommitMessage(commitMessage string) error {
	// Extract the relevant information from the commit message
	lines := strings.Split(commitMessage, "\n")
	if len(lines) < 2 {
		log.Println("Invalid commit message format")
		return fmt.Errorf("invalid commit message format")
	}

	username := lines[0]
	repoBranchLine := lines[1]
	repoBranchParts := strings.Split(repoBranchLine, "[")
	if len(repoBranchParts) != 2 {
		log.Println("Invalid commit message format")
		return fmt.Errorf("invalid commit message format")
	}

	repoBranch := strings.TrimSuffix(repoBranchParts[1], "]")
	repoBranchParts = strings.Split(repoBranch, ":")
	if len(repoBranchParts) != 2 {
		log.Println("Invalid commit message format")
		return fmt.Errorf("invalid commit message format")
	}

	repo := strings.TrimSpace(repoBranchParts[0])
	branch := strings.TrimSpace(repoBranchParts[1])

	commits := lines[2:]

	collection := Database.Collection("commits")

	// Insert the data into MongoDB
	_, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "username", Value: username},
		{Key: "repo", Value: repo},
		{Key: "branch", Value: branch},
		{Key: "commits", Value: commits},
	})
	if err != nil {
		log.Printf("Failed to insert data into MongoDB: %v", err)
		return err
	}

	return nil
}
