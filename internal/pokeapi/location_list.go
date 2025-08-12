package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, exists := c.pokecache.Get(url)
	if exists {
		return unmarshalLocation(val)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.pokecache.Add(url, data)
	return unmarshalLocation(data)
}

func unmarshalLocation(byteData []byte) (RespShallowLocations, error) {
	var locationsResp RespShallowLocations
	if err := json.Unmarshal(byteData, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
