package commands

import (
	"github.com/bwmarrin/discordgo"
)

func handleProjectListInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Extract necessary data from i.Data.Options if needed

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "https://01edu.notion.site/Global-01-Curriculum-50b7d94ac56a429fb3aee19a32248732",
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}
