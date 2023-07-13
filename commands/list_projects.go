package commands

import (
	"github.com/bwmarrin/discordgo"
)

func handleProjectListInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Extract necessary data from i.Data.Options if needed

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Useful link about **01 pedagogy highly recommmend** to checkit out.\nhttps://01edu.notion.site/Content-b00cf9d4179f423f901eeb48e9345c16",
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}
