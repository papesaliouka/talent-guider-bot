package main

import (
	"fmt"
	"golang-discord-bot/commands"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//godotenv.Load(".env")
	err := commands.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	commands.Start()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	// Wait for a termination signal to gracefully close the bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
