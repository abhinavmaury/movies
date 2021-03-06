package requests

import (
	"encoding/json"
	"errors"
	"movies/models/dto"
	"movies/service/auth"
	"net/http"
)

type AddMovieRequest struct {
	Movie dto.MovieInfo `json:"movie"`
}

func (request *AddMovieRequest) Initiate(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		user, err := auth.AuthenticateUser(r.Header.Get("auth_token"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		if user.IsAdmin != 1 {
			err = errors.New("user not authorized")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
	return nil
}
