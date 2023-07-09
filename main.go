package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json: "Token`
	BotPrefix string `json: "BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == BotPrefix+"ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == BotPrefix+"submitproject" {
		err := submitProject(s, m)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error submitting project. Please try again later.")
			fmt.Println("Error submitting project:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Thank you for submitting your project!")
	}

	if m.Content == BotPrefix+"projectlist" {
		projectList, err := getProjectList()
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error retrieving project list. Please try again later.")
			fmt.Println("Error retrieving project list:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Project List:\n"+projectList)
	}

	if m.Content == BotPrefix+"projectdetails" {
		// Extract necessary information from the message, such as project name or identifier
		projectName := "Sample Project"

		projectDetails, err := getProjectDetails(projectName)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error retrieving project details. Please try again later.")
			fmt.Println("Error retrieving project details:", err)
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Project Details:\n"+projectDetails)
	}

	if m.Content == BotPrefix+"solvedproject" {
		incrementSolvedProjects()
		s.ChannelMessageSend(m.ChannelID, "Solved project count incremented.")
	}

	if m.Content == BotPrefix+"help" {
		helpMessage := "Available Commands:\n" +
			BotPrefix + "ping - Check if the bot is responsive.\n" +
			BotPrefix + "submitproject - Submit a project.\n" +
			BotPrefix + "projectlist - Get the list of submitted projects.\n" +
			BotPrefix + "projectdetails - Get the details of a specific project.\n" +
			BotPrefix + "solvedproject - Increment the count of solved projects.\n" +
			BotPrefix + "addproject - Add a project to the list of projects (admin only).\n" +
			BotPrefix + "help - Display this help message."

		s.ChannelMessageSend(m.ChannelID, helpMessage)
	}

	if m.Content == BotPrefix+"addproject" {
		// Check if the user has administrative privileges (you can adjust the condition based on your server's roles/permissions)
		if isAdmin(m.Author.ID) {
			// Extract necessary information from the message, such as project name and description
			projectName := "Sample Project"
			projectDescription := "This is a sample project description."

			err := addProject(projectName, projectDescription)
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

}

func addProject(projectName, projectDescription string) error {
	// Open the file in append mode
	file, err := os.OpenFile("projects.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Format the project information
	projectInfo := fmt.Sprintf("Name: %s\nDescription: %s\n\n", projectName, projectDescription)

	// Write the project information to the file
	_, err = file.WriteString(projectInfo)
	if err != nil {
		return err
	}

	return nil
}

func isAdmin(userID string) bool {
	// Implement your logic to determine if the user is an admin
	// You can use Discord's APIs or check against a list of admin user IDs

	// Sample logic to check if the user is an admin based on their ID
	adminUserIDs := []string{"787648176295378944", "adminUserID2"}
	for _, id := range adminUserIDs {
		if userID == id {
			return true
		}
	}

	return false
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

func incrementSolvedProjects() {
	// Implement your logic to track and increment the count of solved projects
	// You can use a database or other storage mechanism to store the count
	// In this example, we'll assume a simple text file to store the count
	count, err := ioutil.ReadFile("solved_count.txt")
	if err != nil {
		fmt.Println("Error reading solved count:", err)
		return
	}

	solvedCount, err := strconv.Atoi(string(count))
	if err != nil {
		fmt.Println("Error converting solved count:", err)
		return
	}

	solvedCount++
	err = ioutil.WriteFile("solved_count.txt", []byte(strconv.Itoa(solvedCount)), 0644)
	if err != nil {
		fmt.Println("Error writing solved count:", err)
		return
	}
}

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

func main() {
	err := ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	Start()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	// Wait for a termination signal to gracefully close the bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close the Discord session
	defer goBot.Close()

}
