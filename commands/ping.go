package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func handlePingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Calculate the API latency
	apiLatency := s.HeartbeatLatency().Seconds() * 1000

	// Record the start time
	startTime := time.Now()

	// Send the "Pong!" response


	// Calculate the round-trip time (RTT)
	rtt := time.Since(startTime).Seconds() * 1000

	// Send a follow-up message with latency information
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("API Latency: %.2fms\nRound-trip Time (RTT): %.2fms", apiLatency, rtt),
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}
