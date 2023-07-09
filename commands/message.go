package commands

import (
	"github.com/bwmarrin/discordgo"
)

func interactionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {

	switch i.ApplicationCommandData().Name {
	case "ping":
		handlePingInteraction(s, i)
	case "projectlist":
		handleProjectListInteraction(s, i)
	case "help":
		handleHelpInteraction(s, i)
	}

}

func handlePingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}
