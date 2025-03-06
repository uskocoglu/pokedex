package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (PokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationName

	i, ok := c.cache.Get(url)
	if ok {
		//get it from cache
		pokemonResp := PokemonEncounters{}
		err := json.Unmarshal(i, &pokemonResp)
		if err != nil {
			return PokemonEncounters{}, err
		}

		fmt.Println("Got it from Cache!")

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEncounters{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEncounters{}, err
	}

	pokemonResp := PokemonEncounters{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonEncounters{}, err
	}
	// Store the new data in the cache
	c.cache.Add(url, dat)

	return pokemonResp, nil

}