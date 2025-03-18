package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/uskocoglu/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func(*config, ...string) error 
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
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

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 || len(words) > 2{
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(conf, args...)
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
			description: "Get the next page of locations",
			callback:	 commandMapf,
		},
		"mapb": {
			name:		 "mapb",
			description: "Get the previous page of locations",
			callback:	 commandMapb,
		},
		"explore": {
			name:		 "explore <location_name>",
			description: "Explore the pokemons given the location",
			callback:	 commandExplore,
		},
		"catch": {
			name:		 "catch <pokemon_name>",
			description: "Try to catch the pokemon",
			callback:	 commandCatch,
		},
		"inspect": {
			name:		 "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:	 commandInspect,
		},
	}
}