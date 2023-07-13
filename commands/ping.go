package commands

import "github.com/bwmarrin/discordgo"

func handlePingInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}
