package omdb

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	assert := assert.New(t)

	c := &Client{ApiKey: os.Getenv("OMDB_APIKEY")}

	movie, err := c.GetMovie("tt1210166")

	assert.Nil(err, "Error should be nil")
	assert.Equal("Moneyball", movie.Title, "movie title not matching the expected title")

	log.Printf("Movie: %+v", movie)
}

func TestGetSeries(t *testing.T) {
	assert := assert.New(t)

	c := &Client{ApiKey: os.Getenv("OMDB_APIKEY")}

	series, err := c.GetSeries("tt0903747")

	assert.Nil(err, "Error should be nil")
	assert.Equal("Breaking Bad", series.Title, "series title not matching the expected title")

	log.Printf("Series: %+v", series)
}

func TestGetEpisode(t *testing.T) {
	assert := assert.New(t)

	c := &Client{ApiKey: os.Getenv("OMDB_APIKEY")}

	episode, err := c.GetEpisode("tt7029252")

	assert.Nil(err, "Error should be nil")
	assert.Equal("Smoke Signal", episode.Title, "episode title not matching the expected title")

	log.Printf("Episode: %+v", episode)
}
