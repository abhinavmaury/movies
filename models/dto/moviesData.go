package dto

import "movies/models/entity"

type MovieInfo struct {
	Id              int64    `json:"id"`
	PopularityScore float64  `json:"99popularity"`
	Director        string   `json:"director"`
	Genre           []string `json:"genre"`
	ImdbScore       float64  `json:"imdb_score"`
	Name            string   `json:"name"`
}

func MovieInfoDtoToEntity(movie MovieInfo) *entity.Movie {
	return &entity.Movie{
		Id:              movie.Id,
		Name:            movie.Name,
		PopularityScore: movie.PopularityScore,
		Director:        movie.Director,
		ImdbScore:       movie.ImdbScore,
	}
}

func MovieInfoEntityToDto(movie entity.Movie, genre []string) *MovieInfo {
	return &MovieInfo{
		Id:              movie.Id,
		Name:            movie.Name,
		PopularityScore: movie.PopularityScore,
		Director:        movie.Director,
		Genre:           genre,
		ImdbScore:       movie.ImdbScore,
	}
}

func GenreDtoToEntity(movieId int64, genre string) *entity.Genre {
	return &entity.Genre{
		MovieId: movieId,
		Name:    genre,
	}
}
