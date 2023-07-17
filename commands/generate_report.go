package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handleSendReportInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	guildID := i.GuildID

	// Generate reports for all members in the guild
	reports := generateReport(s, guildID)

	// Send reports to all members via private message
	sendReportsToMembers(s, guildID, reports)

	// Send confirmation message
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Reports have been sent to all members.",
		},
	}

	err := s.InteractionRespond(i.Interaction, &response)
	if err != nil {
		log.Printf("Failed to send interaction response: %v", err)
		return
	}
}

func sendReportsToMembers(s *discordgo.Session, guildID string, reports map[string]string) {
	// Get the list of members in the guild
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		log.Printf("Failed to get guild members: %v", err)
		return
	}

	// Loop through each member and send the report as a private message
	for _, member := range members {

		// Skip sending messages to bot users
		if member.User.Bot {
			continue
		}

		// Create a direct message channel with the member
		channel, err := s.UserChannelCreate(member.User.ID)
		if err != nil {
			log.Printf("Failed to create DM channel with user %s: %v", member.User.ID, err)
			continue
		}

		// Get the report for the current member
		report, ok := reports[member.User.ID]
		if !ok {
			log.Printf("Report not found for user %s", member.User.ID)
			continue
		}

		// Send the report as a private message
		_, err = s.ChannelMessageSend(channel.ID, report)
		if err != nil {
			log.Printf("Failed to send report to user %s: %v", member.User.ID, err)
		}
	}
}

func generateReport(s *discordgo.Session, guildID string) map[string]string {
	reports := make(map[string]string)

	// Get the list of members in the guild
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		log.Printf("Failed to get guild members: %v", err)
		return reports
	}

	// Loop through each member and generate their report
	for _, member := range members {
		// Check if the member has any coding sessions
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

	return reports
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
