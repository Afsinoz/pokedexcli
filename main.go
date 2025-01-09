package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"math/rand"

	"github.com/Afsinoz/pokedexcli/internal/pokeapi"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// TODO: next time make it string to Pokemon

func cleanInput(text string) []string {
	splittedString := strings.Fields(text)
	return splittedString
}

func commandExit(config *Config, param string) error {
	defer os.Exit(0)
	fmt.Println("Closing the Pokedex... Goodbye!")
	return fmt.Errorf("Here is the error")
}

func commandHelp(config *Config, param string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	return fmt.Errorf("Here is the error of help")
}

func commandMap(config *Config, param string) error {

	next, previous, places, err := pokeapi.MapGet(config.Next)
	if err != nil {
		fmt.Println(err)
	}
	config.Next = next
	config.Previous = previous
	//	fmt.Println("MAIN FUNCTION This is the next url: ", config.Next)
	//	fmt.Println("MAIN FUNCTION This is the previous url: ", config.Previous)
	for _, place := range places {
		fmt.Println(place)
	}

	return nil
}

func commandMapBack(config *Config, param string) error {
	if config.Previous == "" {
		fmt.Println("This is literally the first location!")
		return nil
	}
	next, previous, places, err := pokeapi.MapGet(config.Previous)
	if err != nil {
		fmt.Println(err)
	}
	config.Next = next
	config.Previous = previous
	for _, place := range places {
		fmt.Println(place)
	}

	return nil
}

func commandExplore(config *Config, param string) error {
	fmt.Println("Exploring ", param)
	areaName := param
	URL := baseURL + "/location-area/"
	locationNames, err := pokeapi.PokeGet(URL, areaName)
	if err != nil {
		return fmt.Errorf("PokeGet problem: ", err)
	}
	fmt.Println("Found Pokemon:")
	for _, locationName := range locationNames {
		fmt.Println("- ", locationName)
	}
	return nil
}

func commandCatch(config *Config, pokemon string) error {
	//TODO: For the next time, you need to get the pokemon information, not just the name:w

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)

	base_experience, err := pokeapi.PokemonInfoGet(pokemon)
	if err != nil {
		fmt.Println("PokeinfoGet function err: ", err)
		return err
	}
	catchingChance := rand.Intn(base_experience)
	if catchingChance >= base_experience/2 {
		fmt.Println(pokemon, "was caught!")
	} else {
		fmt.Println(pokemon, "escaped!")
	}
	//fmt.Println("The exp level of the ", pokemon, " is ", base_experience)
	//fmt.Println("The catching chance is: ", catchingChance)
	return nil

}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, param string) error
}

type Config struct {
	Next     string
	Previous string
}

func main() {

	var config Config

	mapEndPoint := baseURL

	config.Next = mapEndPoint
	config.Previous = ""

	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of the previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "Displays the name of the pokemons in the areas",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catching the pokemon base on Experience level.",
			callback:    commandCatch,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInput = strings.ToLower(userInput)
		userInputList := cleanInput(userInput)
		firstWord := userInputList[0]
		var secondWord string
		if len(userInputList) < 2 {
			secondWord = ""
		} else {
			secondWord = userInputList[1]
		}
		command, ok := supportedCommands[firstWord]
		if ok {
			if command.name == "help" {

				command.callback(&config, secondWord)
				for command := range supportedCommands {
					fmt.Println(supportedCommands[command].name, ":", supportedCommands[command].description)
				}

			} else if err := command.callback(&config, secondWord); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "reading the std input:", err)
	}

}
