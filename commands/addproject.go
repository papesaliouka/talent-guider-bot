package commands

import (
	"fmt"
	"os"
)

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
