package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(cacheData, &locationAreasResp)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("cache missed")

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResp)
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}

func (c *Client) GetLocationAreas(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	cacheData, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit")
		locationArea := LocationArea{}
		err := json.Unmarshal(cacheData, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("cache missed")

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 399 {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if resp.StatusCode > 399 {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}
