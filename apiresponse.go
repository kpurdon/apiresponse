package apiresponse

import (
	"encoding/json"
	"net/http"
)

type Responder struct {
	w http.ResponseWriter
}

func NewResponder(w http.ResponseWriter) *Responder {
	return &Responder{w: w}
}

func (r Responder) write(statusCode int, data interface{}) {
	if data == nil {
		// TODO: something more useful?
		data = struct {
			StatusCode int    `json:"status_code"`
			StatusText string `json:"status_text"`
		}{
			StatusCode: statusCode,
			StatusText: http.StatusText(statusCode),
		}
	}

	b, err := json.Marshal(data)
	if err != nil {
		r.InternalServerError()
		return
	}

	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(statusCode)
	r.w.Write(b)
}

// TODO: generate these based on http.Status*

func (r Responder) OK(data interface{}) {
	r.write(http.StatusOK, data)
}

func (r Responder) InternalServerError() {
	r.write(http.StatusInternalServerError, nil)
}

func (r Responder) NotFound() {
	r.write(http.StatusNotFound, nil)
}

func (r Responder) Created(data interface{}) {
	r.write(http.StatusCreated, data)
}

func (r Responder) NoContent() {
	r.write(http.StatusNoContent, nil)
}

func (r Responder) BadRequest() {
	r.write(http.StatusBadRequest, nil)
}
