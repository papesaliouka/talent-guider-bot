package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func handleShowExerciseInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	levelOption := i.ApplicationCommandData().Options[0]
	level := int(levelOption.IntValue())

	exercises := ShowExercise(level)

	if exercises != nil {
		// Create the options for the select component
		var options []discordgo.SelectMenuOption
		for _, exercise := range exercises {
			option := discordgo.SelectMenuOption{
				Label: exercise,
				Value: exercise,
			}
			options = append(options, option)
		}

		// Create the select component
		selectComponent := discordgo.SelectMenu{
			CustomID:    "select_exercise",
			Placeholder: "Select an exercise",
			Options:     options,
		}

		// Create the response with the select component
		response := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Please select an exercise:",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{&selectComponent},
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

func selectInteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Interaction.Type == discordgo.InteractionMessageComponent {
		data := i.Interaction.MessageComponentData()
		switch data.CustomID {
		case "select_exercise":
			exerciseName := data.Values[0]
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
		log.Printf("Failed to execute wget command: %v", err)
	}

	content, err := ioutil.ReadFile(outputPath)
	if err != nil {
		log.Printf("Failed to read exercise file: %v", err)
	}

	// Create the response with the fetched exercise content
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("```Exercise: %s\n```\n%s\n``````", exerciseName, string(content)),
		},
	}
	s.InteractionRespond(i.Interaction, &response)

	err = exec.Command("rm", outputPath).Run()
	if err != nil {
		log.Printf("Failed to delete exercise file: %v", err)
	}
}

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
