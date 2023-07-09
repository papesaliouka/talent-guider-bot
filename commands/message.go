package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	// Parse the message content to get the command and arguments
	parts := strings.Fields(m.Content)
	if len(parts) < 1 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case BotPrefix + "ping":
		handlePing(s, m, args)
	case BotPrefix + "submitproject":
		handleSubmitProject(s, m, args)
	case BotPrefix + "projectlist":
		handleProjectList(s, m)
	case BotPrefix + "projectdetails":
		handleProjectDetails(s, m, args)
	case BotPrefix + "solvedproject":
		handleSolvedProject(s, m, args)
	case BotPrefix + "addproject":
		handleAddProject(s, m, args)
	case BotPrefix + "help":
		handleHelp(s, m, args)
	}
}

func handlePing(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSend(m.ChannelID, "pong")
}
