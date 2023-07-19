package commands

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleGenerateReportInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guildID := GuildID

	// Generate reports for all members in the guild
	currentTime := time.Now()

	today := time.Now().Format("2006-01-02") // Format today's date as YYYYMMDD
	formattedDate := currentTime.Format("01-2006")
	logfilename := fmt.Sprintf("./data/sessions_logs/%s/%vlog.csv", formattedDate, today)

	reports, err := generateReport(s, guildID, logfilename)
	if err != nil {
		log.Printf("Failed to generate reports: %v", err)
		return
	}

	// Send reports to all members via private message
	generateReportsForMembers(s, i.ChannelID, reports)

	// Send confirmation message
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reports have been sent to all members.",
		},
	}

	err = s.InteractionRespond(i.Interaction, &response)
	if err != nil {
		log.Printf("Failed to send interaction response: %v", err)
		return
	}
}

func generateReportsForMembers(s *discordgo.Session, channelID string, reports map[string]string) {
	for userID, report := range reports {
		// Skip sending messages to bot users
		if isBotUser(s, userID) {
			continue
		}

		// Send the report as a private message
		_, err := s.ChannelMessageSend(channelID, report)
		if err != nil {
			log.Printf("Failed to send report to user %s: %v", userID, err)
		}
	}
}

func isBotUser(s *discordgo.Session, userID string) bool {
	user, err := s.User(userID)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return true // If there's an error, consider it a bot user to be safe
	}
	return user.Bot
}

// readCodingSessionsFromLog reads the log file in CSV format and returns a slice of CodingSessionLog
func readCodingSessionsFromLog(logFilePath string) ([]CodingSessionLog, error) {
	var codingSessions []CodingSessionLog

	file, err := os.Open(logFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read the header row to skip it
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	// Read each row and parse it into CodingSessionLog
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		startTime, err := time.Parse(time.RFC3339, row[1]) // Assuming the start time is in the second column
		if err != nil {
			return nil, err
		}

		endTime, err := time.Parse(time.RFC3339, row[2]) // Assuming the end time is in the third column
		if err != nil {
			return nil, err
		}

		duration, err := time.ParseDuration(row[3]) // Assuming the duration is in the fourth column
		if err != nil {
			return nil, err
		}

		sessionLog := CodingSessionLog{
			UserID:    row[0], // Assuming the user ID is in the first column
			StartTime: startTime,
			EndTime:   endTime,
			Duration:  duration,
		}

		codingSessions = append(codingSessions, sessionLog)
	}

	return codingSessions, nil
}

func generateReport(s *discordgo.Session, guildID string, logFilePath string) (map[string]string, error) {
	reports := make(map[string]string)

	// Get the list of members in the guild
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		log.Printf("Failed to get guild members: %v", err)
		return reports, err
	}

	// Read the coding sessions from the CSV log file
	codingSessions, err := readCodingSessionsFromLog(logFilePath)
	if err != nil {
		log.Printf("Failed to read coding sessions from log file: %v", err)
		return reports, err
	}

	// Loop through each member and generate their report
	for _, member := range members {
		// Check if the member has any coding sessions in the log file
		var totalDuration time.Duration
		for _, session := range codingSessions {
			if session.UserID == member.User.ID {
				totalDuration += session.Duration
			}
		}

		// Generate the report for the current member
		report := generateMemberReport(member.User.Username, totalDuration)

		// Add the report to the map
		reports[member.User.ID] = report
	}

	return reports, nil
}

func generateMemberReport(username string, duration time.Duration) string {
	report := fmt.Sprintf("Member: %s\n", username)

	if duration > 0 {
		report += fmt.Sprintf("Total Coding Duration: %s\n", duration.String())
	} else {
		report += "No coding sessions recorded.\n"
	}

	// Add any additional information or calculations to the report as needed

	return report
}
