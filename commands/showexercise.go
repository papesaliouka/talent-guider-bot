package commands

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"encoding/json"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

type ExerciseData struct {
	Exercises map[string][]string `json:"exercises"`
}

func ShowExercise(level int) []string {
	// Read exercise data from JSON file
	filePath := "data/exercise.json"
	exerciseData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read exercise data:", err)
		return nil
	}

	// Parse JSON data into ExerciseData struct
	var data ExerciseData
	err = json.Unmarshal(exerciseData, &data)
	if err != nil {
		fmt.Println("Failed to parse exercise data:", err)
		return nil
	}

	// Retrieve exercises for the specified level
	levelKey := fmt.Sprintf("%d", level)
	exercises, ok := data.Exercises[levelKey]
	if !ok {
		fmt.Println("No exercises found for level", level)
		return nil
	}

	return exercises
}

func handleShowExerciseInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	levelOption := i.ApplicationCommandData().Options[0]
	level := int(levelOption.IntValue())

	exercises := ShowExercise(level)
	if exercises != nil {
		// Create the content with clickable exercise names
		var content strings.Builder
		content.WriteString(fmt.Sprintf("Exercises for level %d:\n", level))
		for _, exercise := range exercises {
			content.WriteString(fmt.Sprintf("[%s](https://raw.githubusercontent.com/01-edu/public/master/subjects/%s/README.md)\n", exercise, exercise))
		}

		// Create the response with clickable message components
		response := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content.String(),
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							&discordgo.Button{
								Label:    "Fetch Exercise",
								Style:    discordgo.PrimaryButton,
								CustomID: "fetch_exercise",
							},
						},
					},
				},
			},
		}
		s.InteractionRespond(i.Interaction, &response)
	} else {
		response := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("No exercises found for level %d.", level),
			},
		}
		s.InteractionRespond(i.Interaction, &response)
	}
}

func buttonInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Type == discordgo.InteractionMessageComponent {
		switch i.MessageComponentData().CustomID {
		case "fetch_exercise":
			exerciseName := i.MessageComponentData().Values[0]
			fetchExercise(s, i, exerciseName)
		}
	}
}

func fetchExercise(s *discordgo.Session, i *discordgo.InteractionCreate, exerciseName string) {
	fileURL := fmt.Sprintf("https://raw.githubusercontent.com/01-edu/public/master/subjects/%s/README.md", exerciseName)
	outputPath := exerciseName + ".md"

	cmd := exec.Command("wget", "-O", outputPath, fileURL)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to execute wget command: %v", err)
	}

	content, err := ioutil.ReadFile(outputPath)
	if err != nil {
		log.Fatalf("Failed to read exercise file: %v", err)
	}

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Exercise: %s\n```md\n%s\n```", exerciseName, string(content)),
		},
	}
	s.InteractionRespond(i.Interaction, &response)

	err = exec.Command("rm", outputPath).Run()
	if err != nil {
		log.Fatalf("Failed to delete exercise file: %v", err)
	}
}
