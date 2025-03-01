package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// If the data exists in the cache, do not make HTTP request, get it from the cache
	i, ok := c.cache.Get(url)
	if ok {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(i, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		fmt.Println("Got it from Cache!")

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	// Store the new data in the cache
	c.cache.Add(url, dat)

	return locationsResp, nil
}

func (c *Client) ListPokemons(locationArea string) (PokemonEncounters, error) {
	url := baseURL + "/location-area/" + locationArea

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