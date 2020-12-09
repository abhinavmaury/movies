package entity

import (
	"context"
	"movies/base"
)

type Movie struct {
	Id              int64
	Name            string
	PopularityScore float64
	Director        string
	ImdbScore       float64
}

func (movie *Movie) Insert(ctx *context.Context) (int64, error) {
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	query := `INSERT INTO movies (name, popularity_score, director, imdb_score) VALUES (?,?,?,?)`
	row, err := conn.ExecContext(*ctx, query, movie.Name, movie.PopularityScore, movie.Director, movie.ImdbScore)
	if err != nil {
		return 0, err
	}
	id, _ := row.LastInsertId()
	return id, nil
}

func (movie *Movie) Update(ctx *context.Context) error {
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `UPDATE movies set name = ?, popularity_score = ?, director = ?, imdb_score = ? WHERE id = ?`
	_, err = conn.ExecContext(*ctx, query, movie.Name, movie.PopularityScore, movie.Director, movie.ImdbScore, movie.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetMovies(ctx *context.Context) (*[]Movie, error) {
	movies := make([]Movie, 0)
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return &movies, err
	}
	defer conn.Close()
	query := `select id, name, popularity_score, director, imdb_score from movies.movies where is_active = 1`
	rows, err := conn.QueryContext(*ctx, query)
	if err != nil {
		return &movies, err
	}
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Id, &movie.Name, &movie.PopularityScore, &movie.Director, &movie.ImdbScore); err == nil {
			movies = append(movies, movie)
		}
	}
	return &movies, nil
}

func (movie *Movie) Deactivate(ctx *context.Context) error {
	conn, err := base.DB.Conn(*ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	query := `UPDATE movies set is_active = ? WHERE id = ?`
	_, err = conn.ExecContext(*ctx, query, 0, movie.Id)
	if err != nil {
		return err
	}
	return nil
}
