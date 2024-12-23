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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInput = strings.ToLower(userInput)
		userInputList := cleanInput(userInput)
		firstWord := userInputList[0]
		fmt.Println("Your command was:",firstWord)
	}
	if err := scanner.Err() ; err != nil {
		fmt.Fprintln(os.Stderr, "reading the std input:", err)
	}
	
}
