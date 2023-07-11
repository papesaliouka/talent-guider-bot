package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Exercise struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	// Add other exercise properties if needed
}

var previousExercises []string

func handleDailyChallengeInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Get the current date
	currentDate := time.Now()

	// Define the start date for the two-week period
	startDate := time.Date(2023, time.July, 1, 0, 0, 0, 0, time.UTC)

	// Load the last execution date from a file or database
	lastExecutionDate := loadLastExecutionDate()

	// Compare the last execution date with the current date
	if currentDate.Year() == lastExecutionDate.Year() && currentDate.YearDay() == lastExecutionDate.YearDay() {
		// Same day, send the same result as before
		sendPreviousChallengeResult(s, i)
		return
	}

	// Determine the number of days passed since the start of the two-week period
	daysPassed := int(currentDate.Sub(startDate).Hours() / 24)

	// Calculate the level for the current day
	level := (daysPassed % 9) + 1

	// Get the random exercises for the current day
	exercises, err := RandomExerciseChooser(level)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Create the content with the exercise names
	var content strings.Builder
	content.WriteString(fmt.Sprintf("Daily Challenge for Level %d:\n", level))
	for _, exercise := range exercises {
		content.WriteString(fmt.Sprintf("Exercise: %s\n", exercise.Name))
	}

	// Create the response message
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content.String(),
		},
	}

	s.InteractionRespond(i.Interaction, &response)

	storeDailyChallengeResult(content.String())

	// Store the current date as the last execution date
	storeLastExecutionDate(currentDate)

}

func RandomExerciseChooser(level int) (exercises []Exercise, err error) {
	// Read the exercise data from the JSON file
	file, err := ioutil.ReadFile("data/exercise.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise file: %v", err)
	}

	// Unmarshal the JSON data into ExerciseData struct
	var exerciseData ExerciseData
	err = json.Unmarshal(file, &exerciseData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal exercise data: %v", err)
	}

	// Get the exercises for the specified level
	levelExercises, ok := exerciseData.Exercises[strconv.Itoa(level)]
	if !ok {
		return nil, fmt.Errorf("level %d does not exist", level)
	}

	// Shuffle the exercises randomly
	rand.Shuffle(len(levelExercises), func(i, j int) {
		levelExercises[i], levelExercises[j] = levelExercises[j], levelExercises[i]
	})

	// Select three exercises that are not in the previous exercises list
	exercises = make([]Exercise, 0)
	for _, exerciseName := range levelExercises {
		if !contains(previousExercises, exerciseName) {
			exercise := Exercise{
				Name:  exerciseName,
				Level: level,
			}
			exercises = append(exercises, exercise)
			if len(exercises) == 3 {
				break
			}
		}
	}

	return exercises, nil
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Global variable to store the last execution date
var lastExecutionDate time.Time

// Load the last execution date from a file or database
func loadLastExecutionDate() time.Time {
	// Your implementation to load the last execution date from a file or database
	// Example:
	// Read the last execution date from a file
	lastExecutionStr, err := ioutil.ReadFile("last_execution.txt")
	if err != nil {
		// Error handling
		return time.Time{}
	}

	// Parse the last execution date from the file content
	lastExecutionTime, err := time.Parse(time.RFC3339, string(lastExecutionStr))
	if err != nil {
		// Error handling
		return time.Time{}
	}

	return lastExecutionTime
}

// Store the current date as the last execution date
func storeLastExecutionDate(date time.Time) {
	// Update the global variable
	lastExecutionDate = date

	// Your implementation to store the last execution date in a file or database
	// Example:
	// Convert the last execution date to string
	lastExecutionStr := date.Format(time.RFC3339)

	// Write the last execution date to a file
	err := ioutil.WriteFile("last_execution.txt", []byte(lastExecutionStr), 0644)
	if err != nil {
		// Error handling
	}
}

// Send the previous challenge result as the daily challenge
func sendPreviousChallengeResult(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Your implementation to retrieve and send the previous challenge result

	// Example: Replace with your code to retrieve the previous challenge result
	previousChallengeResult := getPreviousChallengeResult()

	// Create the response message with the previous challenge result
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "The daily challenge has already been executed today. Here's the previous result:",
			Embeds: []*discordgo.MessageEmbed{
				{
					Description: previousChallengeResult,
				},
			},
		},
	}

	s.InteractionRespond(i.Interaction, &response)
}

// Store the daily challenge result in a file
func storeDailyChallengeResult(result string) error {
	// Write the daily challenge result to the file
	err := ioutil.WriteFile("daily_challenge.txt", []byte(result), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Retrieve the previous challenge result from a file
func getPreviousChallengeResult() string {
	// Read the content of the previous challenge file
	content, err := ioutil.ReadFile("daily_challenge.txt")
	if err != nil {
		// Error handling
		return ""
	}

	// Convert the content to string and return it as the previous challenge result
	return string(content)
}
