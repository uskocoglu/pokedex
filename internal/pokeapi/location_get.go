package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	i, ok := c.cache.Get(url)
	if ok {
		//get it from cache
		locationResp := Location{}
		err := json.Unmarshal(i, &locationResp)
		if err != nil {
			return Location{}, err
		}

		fmt.Println("Got it from Cache!")

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}
	// Store the new data in the cache
	c.cache.Add(url, dat)

	return locationResp, nil

}