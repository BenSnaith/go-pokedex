package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonData(pokemonName string) (RespPokemonData, error) {
	// get the URL to make the API request
	url := base_url + "/pokemon/" + pokemonName

	// check if the data is already in client pokecache
	if val, ok := c.cache.Get(url); ok {
		pokemonDataResp := RespPokemonData{}
		err := json.Unmarshal(val, &pokemonDataResp)
		if err != nil {
			return RespPokemonData{}, err
		}

		return pokemonDataResp, nil
	}

	// create the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonData{}, err
	}

	// use the client httpClient to make the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonData{}, err
	}

	// defer closing the response body to when the function finished
	defer res.Body.Close()

	// read all the data from the response into []byte
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonData{}, err
	}

	// create a stuct to store our data into
	pokemonDataResp := RespPokemonData{}
	// unmarshal the json response into pokemonDataResp
	err = json.Unmarshal(data, &pokemonDataResp)
	if err != nil {
		return RespPokemonData{}, err
	}

	// add the given response into the client pokecache
	c.cache.Add(url, data)
	// return the response struct
	return pokemonDataResp, nil
}
