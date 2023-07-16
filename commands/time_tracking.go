package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

type CodingSession struct {
	UserID    string        `json:"userID"`
	StartTime time.Time     `json:"startTime"`
	Duration  time.Duration `json:"duration"`
}

var codingSessions []CodingSession

func loadCodingSessionsFromFile() error {
	file, err := os.Open("sessions.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&codingSessions)
	if err != nil {
		return err
	}

	return nil
}

func saveCodingSessionsToFile() error {
	file, err := os.Create("sessions.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(codingSessions)
	if err != nil {
		return err
	}

	return nil
}

func handleStartCodingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Interaction.Member.User.ID

	// Store the coding session in the slice
	codingSessions = append(codingSessions, CodingSession{
		UserID:    userID,
		StartTime: time.Now(),
	})

	// Set a timer for 2 hours
	time.AfterFunc(2*time.Hour, func() {
		// Send a reminder to the user that the coding session has ended
		response := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("<@%s>, your coding session has ended. You can start a new 2-hour session by using the /startCoding command.", userID),
			},
		}
		s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
			Content: &response.Data.Content,
		})
	})

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("<@%s>, your coding session has started. You have 2 hours to code!", userID),
		},
	}

	s.InteractionRespond(i.Interaction, &response)

	// Save the updated coding sessions to the file
	err := saveCodingSessionsToFile()
	if err != nil {
		log.Printf("Failed to save coding sessions: %v", err)
	}
}

func handleEndCodingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Interaction.Member.User.ID

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

	// Calculate the duration of the coding session
	startTime := codingSessions[sessionIndex].StartTime
	duration := time.Since(startTime)

	// Update the coding session with the duration
	codingSessions[sessionIndex].Duration = duration

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("<@%s>, your coding session has ended. You coded for %s. You can start a new 2-hour session by using the /startCoding command.", userID, duration.String()),
		},
	}

	s.InteractionRespond(i.Interaction, &response)

	// Save the updated coding sessions to the file
	err := saveCodingSessionsToFile()
	if err != nil {
		log.Printf("Failed to save coding sessions: %v", err)
	}
}
