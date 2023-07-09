package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
