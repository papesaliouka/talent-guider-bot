package commands

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

func getProjectList() (string, error) {
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

	return string(content), nil
}

func handleProjectList(s *discordgo.Session, m *discordgo.MessageCreate) {
	projectList, err := getProjectList()
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error retrieving project list. Please try again later.")
		fmt.Println("Error retrieving project list:", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, "Project List:\n"+projectList)
	
}
