package main

import (
	"fmt"
	"strings"
)


func cleanInput(text string) []string {
	splittedString := strings.Fields(text)
	return splittedString 
}

func main() {
	text := " hello world  "
	fmt.Println(text)
}
