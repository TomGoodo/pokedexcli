package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tomgoodo/pokedexcli/internal/pokecache"
)

type PokeClient struct {
	httpClient http.Client
	pokecache  pokecache.Cache
}

func NewPokeClient() PokeClient {
	return PokeClient{
		httpClient: http.Client{},
		pokecache:  *pokecache.NewCache(time.Duration(time.Minute)),
	}
}

func (c *PokeClient) ListLocations(url *string) (locations, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area"
	requestUrl := baseURL
	if url != nil {
		requestUrl = *url
	}
	if cachedVal, ok := c.pokecache.Get(requestUrl); ok {
		result := locations{}
		err := json.Unmarshal(cachedVal, &result)
		if err != nil {
			return locations{}, err
		}
		return result, nil
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

func (c *PokeClient) LocationExplore(name string) (location, error) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	requestUrl := baseURL

	requestUrl = requestUrl + name

	if cachedVal, ok := c.pokecache.Get(requestUrl); ok {
		result := location{}
		err := json.Unmarshal(cachedVal, &result)
		if err != nil {
			return location{}, err
		}
		return result, nil
	}
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return location{}, err
	}
	defer resp.Body.Close()

	result := location{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return location{}, err
	}
	return result, nil
}
