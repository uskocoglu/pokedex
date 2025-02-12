package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var url = "https://pokeapi.co/api/v2/location-area"

type cliCommand struct {
	name 		string
	description string
	callback 	func(*config) error 
}

type config struct {
	Next	 *string
	Previous *string
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

var conf = config{
	Next: &url,
	Previous: nil,
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			name:		 "map",
			description: "Displays the names of next 20 location",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays the names of previous 20 location",
			callback:	 commandMapb,
		},
	}
}