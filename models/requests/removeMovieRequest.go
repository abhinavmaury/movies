package requests

import (
	"errors"
	"movies/models/dto"
	"movies/service/auth"
	"net/http"
	"strconv"
)

type RemoveMovieRequest struct {
	Movie dto.MovieInfo `json:"list"`
}

func (request *RemoveMovieRequest) Initiate(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		user, err := auth.AuthenticateUser(r.Header.Get("auth_token"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}
		if user.IsAdmin != 1 {
			err = errors.New("user not authorized")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		id := r.URL.Query().Get("id")
		movieId, err := strconv.ParseInt(id, 10, 64)
		if movieId <= 0 || err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		request.Movie.Id = movieId
	default:
		w.WriteHeader(http.StatusNotFound)
	}
	return nil
}
