package main

import (
	argumentshandler "assignment1/argumentsHandler"
	movies "assignment1/movies"
	restapi "assignment1/restapi"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	arguments := os.Args[1:] // The first element is the path to the command, so we can skip that

	const file string = "movies.db"
	movieDb, err := movies.Connect(file)

	if err != nil {
		log.Fatal(err)
	}

	if len(arguments) > 0 {
		argumentshandler.Start(movieDb, arguments)
	} else {

		restapi.Start(movieDb)
	}

}
