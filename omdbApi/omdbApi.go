package omdbapi

import (
	"assignment1/movies"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Title      string    `json:"Title"`
	Year       string    `json:"Year"`
	Rated      string    `json:"Rated"`
	Released   string    `json:"Released"`
	Runtime    string    `json:"Runtime"`
	Genre      string    `json:"Genre"`
	Director   string    `json:"Director"`
	Writer     string    `json:"Writer"`
	Actors     string    `json:"Actors"`
	Plot       string    `json:"Plot"`
	Language   string    `json:"Language"`
	Country    string    `json:"Country"`
	Awards     string    `json:"Awards"`
	Poster     string    `json:"Poster"`
	Ratings    []Ratings `json:"Ratings"`
	Metascore  string    `json:"Metascore"`
	ImdbRating string    `json:"imdbRating"`
	ImdbVotes  string    `json:"imdbVotes"`
	ImdbID     string    `json:"imdbID"`
	Type       string    `json:"Type"`
	DVD        string    `json:"DVD"`
	BoxOffice  string    `json:"BoxOffice"`
	Production string    `json:"Production"`
	Website    string    `json:"Website"`
	Response   string    `json:"Response"`
}

type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type UrlParams struct {
	I      string `json:"i"`
	Apikey string `json:"apikey"`
}

func Get(movieId string) (Response, error) {
	queryParams := url.Values{
		"i":      {movieId},
		"apikey": {"9b2fb0"},
	}
	response, err := http.Get("http://www.omdbapi.com/?" + queryParams.Encode())
	if err != nil {
		return Response{}, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return Response{}, err
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func Worker(movieDb movies.MovieDb, id int, movieId string, results chan<- int) {
	// fmt.Printf("Worker %d getting summary for %s\n", id, movieId)
	res, err := Get(movieId)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Worker %d received summary for %s\n", id, movieId)

	err = movieDb.UpdateSummary(movieId, res.Plot)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Wrote summary for %s to database\n", movieId)
	results <- 1
}
