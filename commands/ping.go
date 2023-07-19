package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func handlePingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Calculate the API latency
	apiLatency := s.HeartbeatLatency().Seconds() * 1000

	// Send a follow-up message with latency information
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("API Latency: %.2fms", apiLatency),
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}
