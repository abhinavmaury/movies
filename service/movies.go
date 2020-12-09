package service

import (
	"context"
	"movies/models/dto"
	"movies/models/entity"
	"movies/models/requests"
)

func ProcessMovies(movies *requests.MoviesList) error {
	if len(*movies) <= 0 {
		return nil
	}
	for _, movie := range *movies {
		if err := AddNewMovie(&movie); err != nil {
			return err
		}
	}
	return nil
}

func AddNewMovie(movie *dto.MovieInfo) error {
	ctx := context.Background()
	if id, err := dto.MovieInfoDtoToEntity(*movie).
		Insert(&ctx); err != nil {
		return err
	} else if id > 0 {
		for _, genre := range movie.Genre {
			if err = dto.GenreDtoToEntity(id, genre).
				Upsert(&ctx); err != nil {
				return err
			}
		}
	}
	return nil
}

func RemoveMovie(movie *dto.MovieInfo) error {
	ctx := context.Background()
	if err := dto.MovieInfoDtoToEntity(*movie).
		Deactivate(&ctx); err != nil {
		return err
	}
	if err := dto.GenreDtoToEntity(movie.Id, "").
		Deactivate(&ctx); err != nil {
		return err
	}
	return nil
}

func EditMovie(movie *dto.MovieInfo) error {
	ctx := context.Background()
	if err := dto.MovieInfoDtoToEntity(*movie).
		Update(&ctx); err != nil {
		return err
	}
	if err := dto.GenreDtoToEntity(movie.Id, "").
		Deactivate(&ctx); err != nil {
		return err
	}
	for _, genre := range movie.Genre {
		if err := dto.GenreDtoToEntity(movie.Id, genre).
			Upsert(&ctx); err != nil {
			return err
		}
	}
	return nil
}

func ViewMoviesList() (*[]dto.MovieInfo, error) {
	ctx := context.Background()
	moviesList := make([]dto.MovieInfo, 0)
	movies, err := entity.GetMovies(&ctx)
	if err != nil {
		return nil, err
	}
	genres, err := entity.GetGenres(&ctx)
	if err != nil {
		return nil, err
	}
	for _, movie := range *movies {
		moviesList = append(moviesList, *dto.MovieInfoEntityToDto(movie, genres[movie.Id]))
	}
	return &moviesList, nil
}
