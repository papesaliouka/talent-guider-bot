package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleViewProjectInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	projectNameOption := i.ApplicationCommandData().Options[0]
	projectName := projectNameOption.StringValue()

	projectName = strings.ToLower(projectName)

	readmeContent, err := getReadmeContent(projectName)
	if err != nil {
		log.Printf("Failed to fetch readme content: %v", err)
		response := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Failed to fetch readme for project: %s. Maybe the project name you provided is incorrect. Verify if it is correct ", projectName),
			},
		}
		s.InteractionRespond(i.Interaction, response)
		return
	}

	sendMultipleMessages(s, i.Interaction, projectName, readmeContent)
}

func sendMultipleMessages(s *discordgo.Session, interaction *discordgo.Interaction, projectName, content string) {
	const maxContentLength = 1934

	// Split the content into chunks
	chunks := splitContent(content, maxContentLength)

	for i, chunk := range chunks {
		// Create the response message for each chunk
		// response := &discordgo.MessageSend{
		// 	Content: fmt.Sprintf("Readme for project %s (Part %d):\n```%s```", projectName, i+1, chunk),
		// }

		// Send the response message
		_, err := s.ChannelMessageSend(interaction.ChannelID, fmt.Sprintf("```Readme for project %s (Part %d):\n```%s``````", projectName, i+1, strings.ReplaceAll(chunk, "#", "")))
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func splitContent(content string, maxLength int) []string {
	var chunks []string

	for len(content) > maxLength {
		chunk := content[:maxLength]
		content = content[maxLength:]
		chunks = append(chunks, chunk)
	}

	if len(content) > 0 {
		chunks = append(chunks, content)
	}

	return chunks
}

func getReadmeContent(exerciseName string) (string, error) {
	fileURL := fmt.Sprintf("https://raw.githubusercontent.com/01-edu/public/master/subjects/%s/README.md", exerciseName)
	outputPath := exerciseName + ".md"

	cmd := exec.Command("wget", "-O", outputPath, fileURL)
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed to execute wget command: %v", err)
		return "file not found", err
	}

	content, err := ioutil.ReadFile(outputPath)
	if err != nil {
		log.Printf("Failed to read exercise file: %v", err)
		return "file not found", err
	}

	err = exec.Command("rm", outputPath).Run()
	if err != nil {
		log.Printf("Failed to delete exercise file: %v", err)
		return "file not found", err
	}

	return string(content), nil
}
