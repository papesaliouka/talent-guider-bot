package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == BotPrefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == BotPrefix+"submitproject" {
		err := submitProject(s, m)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error submitting project. Please try again later.")
			fmt.Println("Error submitting project:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Thank you for submitting your project!")
	}

	if m.Content == BotPrefix+"projectlist" {
		projectList, err := getProjectList()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error retrieving project list. Please try again later.")
			fmt.Println("Error retrieving project list:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Project List:\n"+projectList)
	}

	if m.Content == BotPrefix+"projectdetails" {
		// Extract necessary information from the message, such as project name or identifier
		projectName := "Sample Project"

		projectDetails, err := getProjectDetails(projectName)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error retrieving project details. Please try again later.")
			fmt.Println("Error retrieving project details:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Project Details:\n"+projectDetails)
	}

	if m.Content == BotPrefix+"solvedproject" {
		incrementSolvedProjects()
		s.ChannelMessageSend(m.ChannelID, "Solved project count incremented.")
	}

	if m.Content == BotPrefix+"help" {
		sendHelpMessage(s, m.ChannelID)
	}

	if m.Content == BotPrefix+"addproject" {
		// Check if the user has administrative privileges (you can adjust the condition based on your server's roles/permissions)
		if isAdmin(m.Author.ID) {
			// Extract necessary information from the message, such as project name and description
			projectName := "Sample Project"
			projectDescription := "This is a sample project description."

			err := addProject(projectName, projectDescription)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Error adding project. Please try again later.")
				fmt.Println("Error adding project:", err)
				return
			}

			s.ChannelMessageSend(m.ChannelID, "Project added successfully!")
		} else {
			s.ChannelMessageSend(m.ChannelID, "You do not have permission to use this command.")
		}
	}

}
