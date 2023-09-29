package main

import (
	"assignment1/connectors"
	"assignment1/handlers"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	arguments := os.Args[1:] // The first element is the path to the command, so we can skip that

	const file string = "movies.db"
	movieDb, err := connectors.Connect(file)

	if err != nil {
		log.Fatal(err)
	}

	if len(arguments) > 0 {
		handlers.StartArguments(movieDb, arguments)
	} else {

		handlers.StartRestApi(movieDb)
	}

}
