package commands

import (
	"encoding/csv"
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

type CodingSessionLog struct {
	UserID    string        `json:"userID"`
	StartTime time.Time     `json:"startTime"`
	EndTime   time.Time     `json:"endTime"`
	Duration  time.Duration `json:"duration"`
}

var codingSessions []CodingSession

const logDateFormat = "01-2006"
const logTimeFormat = "2006_01_02"

func handleStartCodingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Interaction.Member.User.ID
	if userID == "" {
		userID = i.Interaction.User.ID
	}

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
		UserID:    userID,
		StartTime: time.Now(),
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

func saveCodingSessionToLog(userID string, startTime, endTime time.Time, duration time.Duration) error {
	// Convert the session data to CSV format
	row := []string{
		userID,
		startTime.Format(time.RFC3339),
		endTime.Format(time.RFC3339),
		duration.String(),
	}

	// Get the current time
	currentTime := time.Now()

	// Create the folder path
	folderPath := fmt.Sprintf("./data/sessions-logs/%v", currentTime.Format(logDateFormat))

	// Create the necessary folders if they don't exist
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Printf("Failed to create folder: %v", err)
		return err
	}

	// Create the log file path
	logFilePath := fmt.Sprintf("%v/%s_log.csv", folderPath, currentTime.Format(logTimeFormat))

	// Open the log file with the necessary flags and permissions
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}

func handleEndCodingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	userID := i.Interaction.Member.User.ID
	if userID == "" {
		userID = i.Interaction.User.ID
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

	// Update the coding session with the end time and duration
	codingSessions[sessionIndex].Duration = duration

	s.ChannelMessageSend(i.Interaction.ChannelID, fmt.Sprintf("<@%s>, your coding session has ended. You coded for %s. You can start a new 2-hour session by using the /startCoding command.", userID, duration))

	// Remove the coding session from the slice
	// Log the coding session to the log file
	saveCodingSessionToLog(userID, startTime, endTime, duration)

	codingSessions = append(codingSessions[:sessionIndex], codingSessions[sessionIndex+1:]...)
}
