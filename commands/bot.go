package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotId string

func Start() {
	goBot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error reading token",err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println("error reading user",err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(interactionCreateHandler)
	goBot.AddHandler(selectInteractionHandler)

	err = goBot.Open()

	registerSlashCommands(goBot)

	if err != nil {
		fmt.Println("erro openning bot",err.Error())
		return
	}

	fmt.Println("Bot is running")

}
