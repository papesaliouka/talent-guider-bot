package commands

import "github.com/bwmarrin/discordgo"

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	helpMessage := "Available Commands:\n" +
		BotPrefix + "ping - Check if the bot is responsive.\n" +
		BotPrefix + "submitproject [projectName] [projectDescription] - Submit a project.\n" +
		BotPrefix + "projectlist - Get the list of submitted projects.\n" +
		BotPrefix + "projectdetails [projectName] - Get the details of a specific project.\n" +
		BotPrefix + "solvedproject - Increment the count of solved projects.\n" +
		BotPrefix + "addproject [projectName] [projectDescription] - Add a project to the list of projects (admin only).\n" +
		BotPrefix + "help - Display this help message."

	s.ChannelMessageSend(m.ChannelID, helpMessage)
}
