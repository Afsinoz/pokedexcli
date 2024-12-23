package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func cleanInput(text string) []string {
	splittedString := strings.Fields(text)
	return splittedString 
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return fmt.Errorf("Here is the error")
}
func commandHelp() error {
	
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	return fmt.Errorf("Here is the error of help")
}

type cliCommand struct {
	name string
	description string 
	callback func() error
}

func main() {


supportedCommands := map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
}
	scanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInput = strings.ToLower(userInput)
		userInputList := cleanInput(userInput)
		firstWord := userInputList[0]
		command, ok := supportedCommands[firstWord]
		if ok {
			if command.name == "help"{

				command.callback()
				for command := range supportedCommands{
					fmt.Println(supportedCommands[command].name, ":", supportedCommands[command].description)
				}

		}else if err := command.callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
	if err := scanner.Err() ; err != nil {
		fmt.Fprintln(os.Stderr, "reading the std input:", err)
	}
	
}
