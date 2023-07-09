package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotId string

var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(interactionCreateHandler)

	err = goBot.Open()

	registerSlashCommands(goBot)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

}

func registerSlashCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Check if the bot is responsive.",
		},
		{
			Name:        "submitproject",
			Description: "Submit a project.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectname",
					Description: "The name of the project.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectdescription",
					Description: "The description of the project.",
					Required:    true,
				},
			},
		},
		{
			Name:        "projectlist",
			Description: "Get the list of submitted projects.",
		},
		{
			Name:        "projectdetails",
			Description: "Get the details of a specific project.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectname",
					Description: "The name of the project.",
					Required:    true,
				},
			},
		},
		{
			Name:        "solvedproject",
			Description: "Increment the count of solved projects.",
		},
		{
			Name:        "addproject",
			Description: "Add a project to the list of projects.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectname",
					Description: "The name of the project.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "projectdescription",
					Description: "The description of the project.",
					Required:    true,
				},
			},
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
