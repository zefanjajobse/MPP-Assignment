package main

import (
	movies "assignment1/connectors"
	argumentshandler "assignment1/handlers"
	restapi "assignment1/handlers"
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
		argumentshandler.StartArguments(movieDb, arguments)
	} else {
		restapi.StartRestApi(movieDb)
	}

}
