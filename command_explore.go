package main

import "fmt"

func commandExplore(cfg *config, locationArea *string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemons(*locationArea)
	if err != nil {
		return err
	}

	for _, pok := range pokemonResp.PokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}
	return nil
}