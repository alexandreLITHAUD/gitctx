package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PromptForConfirmation(message string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	var confirmed bool
	for {
		fmt.Println(message)
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" || input == "yes" {
			confirmed = true
			break
		} else if input == "n" || input == "no" {
			confirmed = false
			break
		} else {
			fmt.Println("Please answer 'y' or 'n'.")
		}
	}
	return confirmed, nil
}

func PromptForName(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	var name string
	for {
		fmt.Println(message)
		nameInput, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		name = strings.TrimSpace(nameInput)

		if name != "" {
			break
		} else {
			fmt.Println("Name cannot be empty.")
		}
	}
	return name, nil
}
