package handler

import (
	"movies/models/requests"
	"movies/models/responses"
	"movies/service"
	"net/http"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	request := &requests.AddMovieRequest{}
	response := &responses.BaseResponse{}
	if err := request.Initiate(w, r); err != nil {
		return
	}
	if err := service.AddNewMovie(&request.Movie); err != nil {
		response.Fail(w, err)
		return
	}
	response.Success(w)
	return
}

func ProcessMovies(w http.ResponseWriter, r *http.Request) {
	request := &requests.PopulateMoviesRequest{}
	response := &responses.BaseResponse{}
	if err := request.Initiate(w, r); err != nil {
		return
	}
	if err := service.ProcessMovies(&request.MoviesList); err != nil {
		response.Fail(w, err)
		return
	}
	response.Success(w)
	return
}

func RemoveMovie(w http.ResponseWriter, r *http.Request) {
	request := &requests.RemoveMovieRequest{}
	response := &responses.BaseResponse{}
	if err := request.Initiate(w, r); err != nil {
		return
	}
	if err := service.RemoveMovie(&request.Movie); err != nil {
		response.Fail(w, err)
		return
	}
	response.Success(w)
	return
}

func EditMovie(w http.ResponseWriter, r *http.Request) {
	request := &requests.EditMovieRequest{}
	response := &responses.BaseResponse{}
	if err := request.Initiate(w, r); err != nil {
		return
	}
	if err := service.EditMovie(&request.Movie); err != nil {
		response.Fail(w, err)
		return
	}
	response.Success(w)
	return
}

func ViewMoviesList(w http.ResponseWriter, r *http.Request) {
	response := &responses.MoviesListResponse{}

	list, err := service.ViewMoviesList()
	if err != nil {
		response.Fail(w, err)
		return
	}
	response.Success(w, list)
	return
}
