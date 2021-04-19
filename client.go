package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ApiURL = "http://www.omdbapi.com/?apikey=%s&i=%s"
)

type Client struct {
	ApiKey string
}

func (c *Client) GetMovie(imdbId string) (movie *OmdbMovie, err error) {

	response, err := http.Get(fmt.Sprintf(ApiURL, c.ApiKey, imdbId))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(responseData, &movie)

	return
}

func (c *Client) GetSeries(imdbId string) (series *OmdbSeries, err error) {

	response, err := http.Get(fmt.Sprintf(ApiURL, c.ApiKey, imdbId))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(responseData, &series)

	return
}

func (c *Client) GetEpisode(imdbId string) (episode *OmdbEpisode, err error) {

	response, err := http.Get(fmt.Sprintf(ApiURL, c.ApiKey, imdbId))

	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(responseData, &episode)

	return
}
