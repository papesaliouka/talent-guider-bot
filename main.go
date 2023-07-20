package main

import (
	"fmt"
	"golang-discord-bot/commands"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = commands.ReadConfig()

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
