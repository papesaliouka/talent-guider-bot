package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleHelpInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Get all registered slash commands
	commands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		log.Printf("Failed to get slash commands: %s", err)
		return
	}

	// Construct the response message
	var content strings.Builder
	content.WriteString("Available Commands:\n")
	for _, command := range commands {
		content.WriteString(BotPrefix + command.Name + " - " + command.Description + "\n")
	}

	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content.String(),
		},
	}

	// Send the response message
	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Printf("Failed to send help response: %s", err)
	}
}
