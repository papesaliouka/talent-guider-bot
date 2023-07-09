package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func handleHelpInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Construct the response message
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Available Commands:\n" +
				BotPrefix + "projectlist - Get the list of submitted projects.\n" +
				BotPrefix + "help - Display this help message.",
		},
	}

	// Send the response message
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Printf("Failed to send help response: %s", err)
	}
}
