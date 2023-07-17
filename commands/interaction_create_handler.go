package commands

import (
	"github.com/bwmarrin/discordgo"
)

func interactionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Interaction.Type {
	case discordgo.InteractionPing:
		handlePingInteraction(s, i)
	case discordgo.InteractionApplicationCommand:
		switch i.ApplicationCommandData().Name {
		case "ping":
			handlePingInteraction(s, i)
		case "projectlist":
			handleProjectListInteraction(s, i)
		case "viewproject":
			handleViewProjectInteraction(s, i)
		case "viewaudit":
			handleViewAuditInteraction(s, i)
		case "help":
			handleHelpInteraction(s, i)
		case "selectexercise":
			handleShowExerciseInteraction(s, i)
		case "dailychallenge":
			handleDailyChallengeInteraction(s, i)
		case "askchat":
			handleAskChatInteraction(s, i)
		case "startcoding":
			handleStartCodingInteraction(s, i)
		case "endcoding":
			handleEndCodingInteraction(s, i)
		case "sendreport":
			// Handle sendreport interaction
			handleSendReportInteraction(s, i)
		}
	case discordgo.InteractionMessageComponent:
		switch i.MessageComponentData().CustomID {
		case "selectexercise":
			handleShowExerciseInteraction(s, i)
		}
	}
}
