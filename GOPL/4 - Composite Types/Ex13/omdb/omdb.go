package omdb

// APIURL is base link to get resources
const APIURL = "http://omdbapi.com/?apikey=c40244d1&"

// Movie is a struct
type Movie struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Country    string
	Poster     string
	BoxOffice  string
	Production string
	Response   string
}
