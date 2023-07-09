package commands

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func submitProject(s *discordgo.Session, m *discordgo.MessageCreate) error {
	// Extract necessary information from the message, such as project name and description
	projectName := "Sample Project"
	projectDescription := "This is a sample project description."

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
