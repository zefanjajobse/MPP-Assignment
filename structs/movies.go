package structs

import (
	"database/sql"
	"fmt"
)

type Movie struct {
	IMDb_id      string         `json:"imdb_id"`
	Title        string         `json:"title"`
	Rating       float64        `json:"rating"`
	Year         int64          `json:"year"`
	Plot_summary sql.NullString `json:"-"`
}

func (movie Movie) PrintInfo() {
	fmt.Println("IMDb id:", movie.IMDb_id)
	fmt.Println("Title:", movie.Title)
	fmt.Println("Rating:", movie.Rating)
	fmt.Println("Year:", movie.Year)
}
