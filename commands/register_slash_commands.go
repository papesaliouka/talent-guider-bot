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
		}, {
			Name:        "dailychallenge",
			Description: "Get the daily challenge exercises taken from checkpoint.",
		}, {
			Name:        "startcoding",
			Description: "Start coding session",
		},
		{
			Name:        "endcoding",
			Description: "End coding session",
		},
		{
			Name:        "projectlist",
			Description: "Get the list of all 01 projects for the 18month",
		}, {
			Name:        "viewproject",
			Description: "Show the readme of a project in the 01-edu repo only the mandatory projects for now not optionals",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectname",
					Description: "The project name",
					Required:    true,
				},
			},
		}, {
			Name:        "viewaudit",
			Description: "Show the readme of an audit in the 01-edu repo only the mandatory projects.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "auditname",
					Description: "The project name",
					Required:    true,
				},
			},
		},
		{
			Name:        "help",
			Description: "Display the available commands for Talent Guider.",
		}, {
			Name:        "selectexercise",
			Description: "Show the exercise list for a given level in the checpoint.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "level",
					Description: "The level number for the checkpoint exercise list.",
					Required:    true,
				},
			},
		},
	}

	// Register the slash commands globally
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		log.Printf("Failed to register slash commands: %s", err)
	}
}
