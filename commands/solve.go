package commands

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func handleSolvedProject(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	incrementSolvedProjects()
	s.ChannelMessageSend(m.ChannelID, "Solved project count incremented.")
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
