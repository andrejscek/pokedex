package main

import (
	"errors"
)

func commandMapf(cfg *config) ([]string, error) {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return nil, err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	var names []string
	for _, loc := range locationsResp.Results {
		names = append(names, loc.Name)
	}
	return names, nil
}

func commandMapb(cfg *config) ([]string, error) {
	if cfg.prevLocationsURL == nil {
		return nil, errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return nil, err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	var names []string
	for _, loc := range locationResp.Results {
		names = append(names, loc.Name)
	}
	return names, nil
}
