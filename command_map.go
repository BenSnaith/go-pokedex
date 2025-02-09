package main

import (
	"errors"
	"fmt"
)

func commandMapf(conf *config, args ...string) error {
	locationsResp, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(conf *config, args ...string) error {
	if conf.prevLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
