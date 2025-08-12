package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(locationName string) (LocationResp, error) {
	url := baseURL + "/location-area"
	if locationName != "" {
		url += "/" + locationName
	}

	val, exists := c.pokecache.Get(url)
	if exists {
		return unmarshalEncounters(val)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResp{}, fmt.Errorf("Could not create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResp{}, fmt.Errorf("Could not get response: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResp{}, fmt.Errorf("Could not read the response body: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		fmt.Println("Location not found")
		return LocationResp{}, nil
	}

	if resp.StatusCode != http.StatusOK {
		return LocationResp{}, fmt.Errorf("server returned %d: %s", resp.StatusCode, string(data))
	}

	c.pokecache.Add(url, data)
	return unmarshalEncounters(data)
}

func unmarshalEncounters(byteData []byte) (LocationResp, error) {
	var locationResp LocationResp
	if err := json.Unmarshal(byteData, &locationResp); err != nil {
		return LocationResp{}, fmt.Errorf("Could not unmarshal the response: %w", err)
	}

	return locationResp, nil
}
