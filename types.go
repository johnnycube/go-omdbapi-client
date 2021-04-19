package omdb

// OmdbMovie struct store all data of a movie.
type OmdbMovie struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
}

// OmdbSeries struct store all data of a series.
type OmdbSeries struct {
	Title        string   `json:"Title"`
	Year         string   `json:"Year"`
	Rated        string   `json:"Rated"`
	Released     string   `json:"Released"`
	Runtime      string   `json:"Runtime"`
	Genre        string   `json:"Genre"`
	Director     string   `json:"Director"`
	Writer       string   `json:"Writer"`
	Actors       string   `json:"Actors"`
	Plot         string   `json:"Plot"`
	Language     string   `json:"Language"`
	Country      string   `json:"Country"`
	Awards       string   `json:"Awards"`
	Poster       string   `json:"Poster"`
	Ratings      []Rating `json:"Ratings"`
	Metascore    string   `json:"Metascore"`
	ImdbRating   string   `json:"imdbRating"`
	ImdbVotes    string   `json:"imdbVotes"`
	ImdbID       string   `json:"imdbID"`
	Type         string   `json:"Type"`
	TotalSeasons string   `json:"totalSeasons"`
	Response     string   `json:"Response"`
}

// OmdbEpisode struct store all data of an episode.
type OmdbEpisode struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Season     string   `json:"Season"`
	Episode    string   `json:"Episode"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	SeriesId   string   `json:"seriesID"`
	Type       string   `json:"Type"`
	Response   string   `json:"Response"`
}

// Rating struct store different kinds of ratings.
type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
