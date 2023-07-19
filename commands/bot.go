package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotId string

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
	goBot.AddHandler(selectInteractionHandler)

	err = goBot.Open()

	registerSlashCommands(goBot)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

}
