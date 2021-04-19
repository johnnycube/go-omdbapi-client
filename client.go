// Package omdb is a client in Go for the omdbapi.com json api. With this
// package you can make http requests to https://www.omdbapi.com/.
//
// You can install this package with:
//
//   $ go get github.com/johnnycube/go-omdbapi-client
//
// See also
//
// https://www.omdbapi.com/ for omdbap documentation.
package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ApiURL = "https://www.omdbapi.com/?apikey=%s&i=%s"
)

// Client does the work of perform the REST requests to the omdbapi endpoints.
type Client struct {
	ApiKey string
}

// GetMovie retrieves a movie item based on the provided imdbID.
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

// GetSeries retrieves a series item based on the provided imdbID.
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

// GetEpisode retrieves an episode item based on the provided imdbID.
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
