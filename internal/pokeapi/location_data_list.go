package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationData(locationName string) (RespLocationData, error) {
	url := base_url + "/location-area/" + locationName

	// check if the data is already in client pokecache
	if val, ok := c.cache.Get(url); ok {
		locationDataResp := RespLocationData{}
		err := json.Unmarshal(val, &locationDataResp)
		if err != nil {
			return RespLocationData{}, err
		}

		return locationDataResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationData{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationData{}, err
	}

	locationDataResp := RespLocationData{}
	err = json.Unmarshal(data, &locationDataResp)
	if err != nil {
		return RespLocationData{}, err
	}

	c.cache.Add(url, data)

	return locationDataResp, nil
}
