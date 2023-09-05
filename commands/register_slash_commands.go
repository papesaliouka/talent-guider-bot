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
			Description: "Start coding session", Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "subject",
					Description: "project-name/documentation/research/etc",
					Required:    true,
				},
			},
		},
		{
			Name:        "endcoding",
			Description: "End coding session",
		},
		{
			Name:        "projectlist",
			Description: "Get the list of all 01 projects for the 18 month",
		}, {
			Name:        "viewproject",
			Description: "Github link of a project in the 01-edu repo",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectname",
					Description: "The project name",
					Required:    true,
				},
			},
		},{
			Name:        "help",
			Description: "Display the available commands for Talent Guider.",
		}, {
			Name:        "selectexercise",
			Description: "Show the exercise list for a given level in the checkpoint.",
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
