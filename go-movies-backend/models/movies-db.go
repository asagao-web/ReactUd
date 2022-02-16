package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get retruns one movie and error, if any
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, Rating, runtime, mpaa_rating, created_at, updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// get the genres
	query = `select g.id, g.genre_name, g.created_at, g.updated_at, mg.id, mg.movie_id, mg.genre_id, mg.created_at, mg.updated_at from genres as g inner join movies_genres as mg on g.id = mg.genre_id where mg.movie_id = $1`
	rows, err := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var genre Genre
	var moviegenre MovieGenre
	genres := make(map[int]string)
	for rows.Next() {
		err = rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
			&moviegenre.ID,
			&moviegenre.MovieID,
			&moviegenre.GenreID,
			&moviegenre.CreatedAt,
			&moviegenre.UpdatedAt,
		)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok,", genres)
		}
		genres[moviegenre.ID] = genre.GenreName
		moviegenre.Genre = genre
		// genres = append(genres, moviegenre)

	}
	movie.MovieGenre = genres

	return &movie, nil
}

// All return all movies and error, if any
func (m *DBModel) All(genre ...int) ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genre) > 0 {
		where = fmt.Sprintf("where id in (select  movie_id from movies_genres where genre_id = %d)", genre[0])
	}

	query := fmt.Sprintf(`select id, title, description, year, release_date, Rating, runtime, mpaa_rating, created_at, updated_at from movies %s order by title`, where)

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.Rating,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		query = `select g.id, g.genre_name, g.created_at, g.updated_at, mg.id, mg.movie_id, mg.genre_id, mg.created_at, mg.updated_at from genres as g inner join movies_genres as mg on g.id = mg.genre_id where mg.movie_id = $1`
		rows, err := m.DB.QueryContext(ctx, query, movie.ID)
		defer rows.Close()
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		var genre Genre
		var moviegenre MovieGenre
		genres := make(map[int]string)
		for rows.Next() {
			err = rows.Scan(
				&genre.ID,
				&genre.GenreName,
				&genre.CreatedAt,
				&genre.UpdatedAt,
				&moviegenre.ID,
				&moviegenre.MovieID,
				&moviegenre.GenreID,
				&moviegenre.CreatedAt,
				&moviegenre.UpdatedAt,
			)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("ok,", genres)
			}
			genres[moviegenre.ID] = genre.GenreName
		}
		movie.MovieGenre = genres
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (m *DBModel) GenresAll() ([]*Genre, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, genre_name, created_at, updated_at from genres order by genre_name`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var genres []*Genre

	for rows.Next() {
		var g Genre
		err := rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &g)
	}
	return genres, nil
}
