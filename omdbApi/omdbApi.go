package omdbapi

import (
	"encoding/json"
	"io"
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
