package responses

import (
	"encoding/json"
	"movies/models/dto"
	"net/http"
)

type MoviesListResponse struct {
	BaseResponse
	Movies *[]dto.MovieInfo `json:"movies"`
}

func (r *MoviesListResponse) Success(w http.ResponseWriter, list *[]dto.MovieInfo) {
	r.Status = true
	r.Message = "success"
	r.Movies = list
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(r)
	_, _ = w.Write(data)
}
