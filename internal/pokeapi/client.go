package pokeapi

import (
	"encoding/json"
	"net/http"
)

type PokeClient struct {
	httpClient http.Client
}

func NewPokeClient() PokeClient {
	return PokeClient{
		httpClient: http.Client{},
	}
}

func (c *PokeClient) ListLocations(url *string) (locations, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	requestUrl := baseURL
	if url != nil {
		requestUrl = *url
	}
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locations{}, err
	}
	defer resp.Body.Close()

	result := locations{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return locations{}, err
	}
	return result, nil
}
