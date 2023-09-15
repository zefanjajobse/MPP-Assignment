package argumentshandler

import (
	"assignment1/movies"
	"flag"
	"fmt"
	"log"
)

func Start(movieDb movies.MovieDb, arguments []string) {

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	addImdbId := addCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")
	addTitle := addCommand.String("title", "Carmencita", "The movie's or series' title")
	addYear := addCommand.Int("year", 1894, "The movie's or series' year of release")
	addImdbRating := addCommand.Float64("rating", 5.7, "The movie's or series' rating on IMDb")

	detailsCommand := flag.NewFlagSet("details", flag.ExitOnError)
	detailsImdbId := detailsCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")

	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteImdbId := deleteCommand.String("imdbid", "tt0000001", "The IMDb ID of a movie or series")

	switch arguments[0] {
	case "add":
		addCommand.Parse(arguments[1:])
		movie, err := movieDb.Insert(movies.Movie{IMDb_id: *addImdbId, Title: *addTitle, Rating: *addImdbRating, Year: int64(*addYear)})

		checkError(err)
		movie.PrintInfo()
	case "list":
		res, err := movieDb.AllTitles()
		checkError(err)
		for _, value := range res {
			fmt.Println(value)
		}
	case "details":
		detailsCommand.Parse(arguments[1:])
		res, err := movieDb.FindOne(*detailsImdbId)
		checkError(err)
		res.PrintInfo()
	case "delete":
		deleteCommand.Parse(arguments[1:])
		movieDb.Delete(*deleteImdbId)
	}

	movieDb.Conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
