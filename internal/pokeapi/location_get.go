package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (RespLocation, error) {

	if len(name) == 0 {
		return RespLocation{}, fmt.Errorf("name is required")
	}

	url := baseURL + "/location-area/" + name

	cachedData, found := c.cache.Get(url)
	if found {
		locationResp := RespLocation{}
		err := json.Unmarshal(cachedData, &locationResp)
		if err != nil {
			return RespLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return RespLocation{}, fmt.Errorf("location not found")
	}

	if resp.StatusCode != http.StatusOK {
		return RespLocation{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}
	locationResp := RespLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
