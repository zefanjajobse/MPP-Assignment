package movies

import (
	"database/sql"
	"fmt"
)

type Movie struct {
	IMDb_id string
	Title   string
	Rating  float64
	Year    int64
}

type MovieDb struct {
	Conn *sql.DB
}

func Connect(file string) (MovieDb, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return MovieDb{}, err
	}
	return MovieDb{Conn: db}, nil
}

func (c *MovieDb) AllTitles() ([]string, error) {
	res := []string{}
	rows, err := c.Conn.Query("SELECT Title FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		res = append(res, name)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (c *MovieDb) FindOne(id string) (Movie, error) {
	row := c.Conn.QueryRow("SELECT * FROM movies WHERE IMDb_id=?", id, 4)
	movie := Movie{}
	err := row.Scan(&movie.IMDb_id, &movie.Title, &movie.Rating, &movie.Year)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

func (c *MovieDb) Insert(movie Movie) (Movie, error) {
	res, err := c.Conn.Exec("INSERT INTO movies VALUES(?,?,?,?);", movie.IMDb_id, movie.Title, movie.Rating, movie.Year)
	if err != nil {
		return Movie{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}

func (c *MovieDb) Delete(id string) error {
	res, err := c.Conn.Exec("DELETE FROM movies WHERE IMDb_id=?", id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Movie deleted")
	return nil
}

func (movie Movie) PrintInfo() {
	fmt.Println("IMDb id:", movie.IMDb_id)
	fmt.Println("Title:", movie.Title)
	fmt.Println("Rating:", movie.Rating)
	fmt.Println("Year:", movie.Year)
}
