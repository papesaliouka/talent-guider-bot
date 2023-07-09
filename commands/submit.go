package commands

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func handleSubmitProject(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	err := submitProject(s, m, args)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error submitting project. Please try again later.")
		fmt.Println("Error submitting project:", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Thank you for submitting your project!")
}

func submitProject(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	// Extract necessary information from the message, such as project name and description
	projectName := args[0]
	projectDescription := args[1]

	// Format the project information
	projectInfo := fmt.Sprintf("Name: %s\nDescription: %s\n\n", projectName, projectDescription)

	// Open the file in append mode
	file, err := os.OpenFile("projects.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the project information to the file
	_, err = file.WriteString(projectInfo)
	if err != nil {
		return err
	}

	return nil
}
