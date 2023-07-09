package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleProjectDetails(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Extract necessary information from the message arguments, such as project name or identifier
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, "Please provide a project name.")
		return
	}

	projectName := args[0]

	projectDetails, err := getProjectDetails(projectName)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error retrieving project details. Please try again later.")
		fmt.Println("Error retrieving project details:", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Project Details:\n"+projectDetails)
}

func getProjectDetails(projectName string) (string, error) {
	// Read the project list file
	file, err := os.Open("projects.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the contents of the file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Search for the project details based on the project name or identifier
	// You can customize this part to extract the specific project details you want to display
	// For example, you can use regular expressions or parsing techniques
	// Assuming each project is separated by a double newline (\n\n)
	projects := strings.Split(string(content), "\n\n")
	for _, project := range projects {
		lines := strings.Split(project, "\n")
		if len(lines) > 0 && lines[0] == "Name: "+projectName {
			return project, nil
		}
	}

	return "", fmt.Errorf("Project not found")
}
