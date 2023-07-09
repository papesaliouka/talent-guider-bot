package commands

import "github.com/bwmarrin/discordgo"

func sendHelpMessage(s *discordgo.Session, channelID string) {
	helpMessage := "Available Commands:\n" +
		BotPrefix + "ping - Check if the bot is responsive.\n" +
		BotPrefix + "submitproject - Submit a project.\n" +
		BotPrefix + "projectlist - Get the list of submitted projects.\n" +
		BotPrefix + "projectdetails [projectName] - Get the details of a specific project.\n" +
		BotPrefix + "solvedproject - Increment the count of solved projects.\n" +
		BotPrefix + "addproject [projectName] [projectDescription] - Add a project to the list of projects (admin only).\n" +
		BotPrefix + "help - Display this help message."

	s.ChannelMessageSend(channelID, helpMessage)
}
