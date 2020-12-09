package entity

import (
	"context"
	"movies/base"
)

type Genre struct {
	MovieId int64
	Name    string
}

func (genre Genre) Upsert(ctx *context.Context) error {
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `INSERT INTO genre (movie_id, name) VALUES (?,?) ON DUPLICATE KEY UPDATE is_active = 1`
	_, err = conn.ExecContext(*ctx, query, genre.MovieId, genre.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetGenres(ctx *context.Context) (map[int64][]string, error) {
	genresMap := make(map[int64][]string)
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return genresMap, err
	}
	defer conn.Close()
	query := `SELECT movie_id, name from genre where is_active = 1`
	rows, err := conn.QueryContext(*ctx, query)
	if err != nil {
		return genresMap, err
	}
	var genreList []string
	for rows.Next() {
		var genre Genre
		if err := rows.Scan(&genre.MovieId, &genre.Name); err == nil {
			if genres, ok := genresMap[genre.MovieId]; ok {
				genres = append(genres, genre.Name)
				genresMap[genre.MovieId] = genres
			} else {
				genreList = append(genreList, genre.Name)
				genresMap[genre.MovieId] = genreList
				genreList = nil
			}
		}
	}
	return genresMap, nil
}

func (genre Genre) Deactivate(ctx *context.Context) error {
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `UPDATE genre set is_active = ? WHERE movie_id = ?`
	_, err = conn.ExecContext(*ctx, query, 0, genre.MovieId)
	if err != nil {
		return err
	}
	return nil
}