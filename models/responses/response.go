package responses

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
}

func (r *BaseResponse) Fail(w http.ResponseWriter, err error) {
	r.Status = false
	r.Message = err.Error()
	data, _ := json.Marshal(r)
	_, _ = w.Write(data)
}

func (r *BaseResponse) Success(w http.ResponseWriter) {
	r.Status = true
	r.Message = "success"
	data, _ := json.Marshal(r)
	_, _ = w.Write(data)
}
