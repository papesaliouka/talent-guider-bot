package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func registerSlashCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Check if the bot is responsive.",
		},
		{
			Name:        "projectlist",
			Description: "Get the list of submitted projects.",
		},
		{
			Name:        "help",
			Description: "Display the available commands.",
		},
	}

	// Register the slash commands globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		log.Printf("Failed to register slash commands: %s", err)
	}
}
