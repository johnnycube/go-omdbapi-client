[![Test](https://github.com/johnnycube/go-omdbapi-client/actions/workflows/test.yml/badge.svg)](https://github.com/johnnycube/go-omdbapi-client/actions/workflows/test.yml)
[![GoDoc](https://godoc.org/github.com/johnnycube/go-omdbapi-client?status.svg)](https://godoc.org/github.com/johnnycube/go-omdbapi-client)

# go-omdbapi-client

omdbapi.com client written in go

The omdb api documentation [can be found here](http://www.omdbapi.com/).

## Installation

Install it yourself as:

    $ go get github.com/johnnycube/go-omdbapi-client

(optional) To run unit tests:

    $ cd $GOPATH/src/github.com/johnnycube/go-omdbapi-client
    $ OMDB_APIKEY=your_apikey go test -v

## Usage

First of all you need to get your API key, User key and User name:

* Generate an API key on http://www.omdbapi.com/apikey.aspx

```Go
package main

import (
	"fmt"

	"github.com/johnnycube/go-omdbapi-client"
)

func main() {
	c := omdb.Client{ApiKey: "YOUR API KEY"}

	movie, err := c.GetMovie("tt1210166")
	if err != nil {
		panic(err)
	}

	// Print the title of the movie
	fmt.Printf("Movie: %s", movie.Title)
	// Output: Movie: Moneyball
}
```

The complete __documentation__ can be found [here](https://godoc.org/github.com/johnnycube/go-omdbapi-client).

## Available endpoints

The following functions are available at the moment:
* GetMovie
* GetSeries
* GetEpisode

Searching is currently not implemented

## Contributing

Bug reports and pull requests are welcome on GitHub at github.com/johnnycube/go-omdbapi-client.

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
