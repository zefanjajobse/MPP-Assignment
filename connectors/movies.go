package connectors

import (
	"assignment1/structs"
	"database/sql"
	"fmt"
)

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

func (c *MovieDb) AllIds() ([]string, error) {
	res := []string{}
	rows, err := c.Conn.Query("SELECT IMDb_id FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var IMDb_id string
		err := rows.Scan(&IMDb_id)
		res = append(res, IMDb_id)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (c *MovieDb) All(pagination structs.Pagination) ([]structs.Movie, error) {
	res := []structs.Movie{}
	query_string := "SELECT * FROM movies ORDER BY IMDb_id"
	if pagination.Limit > 0 || pagination.Offset > 0 {
		// + " OFFSET 0 ROWS FETCH NEXT 5 ROWS ONLY"
		query_string += " LIMIT " + fmt.Sprint(pagination.Limit) + " OFFSET " + fmt.Sprint(pagination.Offset)
	}
	rows, err := c.Conn.Query(query_string)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var movie structs.Movie
		err := rows.Scan(&movie.IMDb_id, &movie.Title, &movie.Rating, &movie.Year, &movie.Plot_summary)
		res = append(res, movie)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (c *MovieDb) Count() (int, error) {
	row := c.Conn.QueryRow("SELECT COUNT(*) FROM movies")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *MovieDb) FindOne(id string) (structs.Movie, error) {
	row := c.Conn.QueryRow("SELECT * FROM movies WHERE IMDb_id=?", id, 5)
	movie := structs.Movie{}
	err := row.Scan(&movie.IMDb_id, &movie.Title, &movie.Rating, &movie.Year, &movie.Plot_summary)
	if err != nil {
		return structs.Movie{}, err
	}
	return movie, nil
}

func (c *MovieDb) Insert(movie structs.Movie) (structs.Movie, error) {
	res, err := c.Conn.Exec("INSERT INTO movies VALUES(?,?,?,?,?);", movie.IMDb_id, movie.Title, movie.Rating, movie.Year, nil)
	if err != nil {
		return structs.Movie{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return structs.Movie{}, err
	}

	return movie, nil
}

func (c *MovieDb) UpdateSummary(id string, summary string) error {
	res, err := c.Conn.Exec("UPDATE movies SET Plot_summary=? WHERE IMDb_id=?", summary, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
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
