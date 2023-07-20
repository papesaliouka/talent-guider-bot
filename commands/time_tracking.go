package commands

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CodingSession struct {
	UserID      string        `json:"userID"`
	StartTime   time.Time     `json:"startTime"`
	Duration    time.Duration `json:"duration"`
	SubjectName string        `json:"subjectName"`
}

type CodingSessionLog struct {
	UserID      string        `json:"userID"`
	StartTime   time.Time     `json:"startTime"`
	EndTime     time.Time     `json:"endTime"`
	Duration    time.Duration `json:"duration"`
	SubjectName string        `json:"subjectName"`
}

var codingSessions []CodingSession

func handleStartCodingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Interaction.Member.User.ID
	if userID == "" {
		userID = i.Interaction.User.ID
	}

	subjectNameOption := i.ApplicationCommandData().Options[0]
	subjectName := subjectNameOption.StringValue()

	subjectName = strings.ToLower(subjectName)

	// Check if there is an active coding session for the user
	for _, session := range codingSessions {
		if session.UserID == userID {
			response := discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("<@%s>, you already have an active coding session. You cannot start a new session until the current session ends.", userID),
				},
			}
			s.InteractionRespond(i.Interaction, &response)
			return
		}
	}

	// Store the coding session in the slice
	codingSessions = append(codingSessions, CodingSession{
		UserID:      userID,
		StartTime:   time.Now(),
		SubjectName: subjectName,
	})

	// Set a timer for 2 hours
	time.AfterFunc(2*time.Hour, func() {
		handleEndCodingInteraction(s, i)
	})

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("<@%s>, your coding session has started. You have 2 hours to code!", userID),
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}

func saveCodingSessionToLog(userID string, startTime, endTime time.Time, duration time.Duration, subjectName, username string, collection *mongo.Collection) error {

	// Insert the data into MongoDB
	_, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "userID", Value: userID},
		{Key: "username", Value: username},
		{Key: "startTime", Value: startTime},
		{Key: "endTime", Value: endTime},
		{Key: "duration", Value: duration},
		{Key: "subjectName", Value: subjectName},
	})
	if err != nil {
		log.Printf("Failed to insert data into MongoDB: %v", err)
		return err
	}

	return nil
}

func handleEndCodingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Interaction.Member.User.ID
	username := i.Interaction.Member.User.Username
	if userID == "" {
		userID = i.Interaction.User.ID
		username = i.Interaction.User.Username
	}

	// Find the coding session for the user
	var sessionIndex = -1
	for index, session := range codingSessions {
		if session.UserID == userID {
			sessionIndex = index
			break
		}
	}

	if sessionIndex == -1 {
		// Coding session not found for the user
		response := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "You haven't started a coding session. Use the /startCoding command to begin.",
			},
		}
		s.InteractionRespond(i.Interaction, &response)
		return
	}

	// Calculate the duration and end time of the coding session
	startTime := codingSessions[sessionIndex].StartTime
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	subjectName := codingSessions[sessionIndex].SubjectName

	// Update the coding session with the end time and duration
	codingSessions[sessionIndex].Duration = duration

	s.ChannelMessageSend(i.Interaction.ChannelID, fmt.Sprintf("<@%s>, your coding session has ended. You coded for %s. You can start a new 2-hour session by using the /startCoding command.", userID, duration))

	// Remove the coding session from the slice
	// Log the coding session to the log file
	saveCodingSessionToLog(userID, startTime, endTime, duration, subjectName, username, Collection)

	codingSessions = append(codingSessions[:sessionIndex], codingSessions[sessionIndex+1:]...)
}
