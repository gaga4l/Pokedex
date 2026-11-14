package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeApiResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokeApiResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokeApiResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeApiResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeApiResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return PokeApiResponse{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeApiResponse{}, err
	}

	locationsResp := PokeApiResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokeApiResponse{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
