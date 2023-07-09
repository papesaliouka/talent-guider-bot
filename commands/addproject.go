package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleAddProject(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Check if the user has administrative privileges (you can adjust the condition based on your server's roles/permissions)
	if isAdmin(m.Author.ID) {
		// Extract necessary information from the message arguments, such as project name and description
		if len(args) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Please provide a project name and description.")
			return
		}

		projectName := args[0]
		otherArgs := strings.Split(strings.Join(args[1:], " "), "_")
		fmt.Println(otherArgs)
		projectDescription,projectRequirements, projectAdditionalInfos := otherArgs[1],otherArgs[2],otherArgs[3] 

		err := addProject(projectName, projectDescription,projectRequirements,projectAdditionalInfos)
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

func addProject(projectName, projectDescription, projectRequirements, projectAdditionalInfos string) error {
	// Open the file in append mode
	file, err := os.OpenFile("projects.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Format the project information
	projectInfo := fmt.Sprintf("**Name**: %s\n**Description**: %s\n**Requirements**: %s\n**Additional Infos**: %s\n\n", projectName, projectDescription, projectRequirements, projectAdditionalInfos)

	fmt.Println(projectInfo)

	// Write the project information to the file
	_, err = file.WriteString(projectInfo)
	if err != nil {
		return err
	}

	return nil
}
