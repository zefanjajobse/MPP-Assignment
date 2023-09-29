package connectors

import (
	"assignment1/structs"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Get(movieId string) (structs.Response, error) {
	queryParams := url.Values{
		"i":      {movieId},
		"apikey": {"9b2fb0"},
	}
	response, err := http.Get("http://www.omdbapi.com/?" + queryParams.Encode())
	if err != nil {
		return structs.Response{}, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return structs.Response{}, err
	}

	var responseObject structs.Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func Worker(movieDb MovieDb, id int, movieId string, results chan<- int) {
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
