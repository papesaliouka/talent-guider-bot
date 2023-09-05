package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func handleViewAuditInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	projectNameOption := i.ApplicationCommandData().Options[0]
	projectName := projectNameOption.StringValue()

	projectName = strings.ToLower(projectName)
    projectUrl:=fmt.Sprintf("https://github.com/01-edu/public/tree/master/subjects/%s/audit", projectName)
    response := &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: fmt.Sprintf("Audit for project: %s \n %s", projectName, projectUrl),
        },
    }
    s.InteractionRespond(i.Interaction, response)



	//readmeContent, err := getAuditReadmeContent(projectName)
	//if err != nil {
	//	log.Printf("Failed to fetch readme content: %v", err)
	//	response := &discordgo.InteractionResponse{
	//		Type: discordgo.InteractionResponseChannelMessageWithSource,
	//		Data: &discordgo.InteractionResponseData{
	//			Content: fmt.Sprintf("Failed to fetch readme for project: %s. Maybe the project name you provided is incorrect. Verify if it is correct ", projectName),
	//		},
	//	}
	//	s.InteractionRespond(i.Interaction, response)
	//	return
	//}

    //sendMultipleMessages(s, i.Interaction, projectName, readmeContent)
}

func getAuditReadmeContent(exerciseName string) (string, error) {
	fileURL := fmt.Sprintf("https://raw.githubusercontent.com/01-edu/public/master/subjects/%s/audit/README.md", exerciseName)
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
