package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(conf *config) error {
	res, err := http.Get(*conf.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}

	for i := 0; i < len(locationAreas.Results); i++ {
		fmt.Println(locationAreas.Results[i].Name)
	}
	conf.Next = locationAreas.Next
	conf.Previous = locationAreas.Previous
	return nil
}