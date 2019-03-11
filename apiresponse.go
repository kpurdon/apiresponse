//go:generate ./generate_helpers.sh

// Package apiresponse provides a simple helper for writing JSON formatted responses in an HTTP API.
package apiresponse

import (
	"encoding/json"
	"net/http"
)

// GenericError is returned when no specific data is included (via responder.WithData()).
// TODO: this type could potentially be more useful?
type GenericError struct {
	StatusCode int    `json:"status_code"`
	StatusText string `json:"status_text"`
}

// Responder is the base type which all helper methods hang on. It should be initialized with the
// NewResponder method.
type Responder struct {
	data interface{}
	http.ResponseWriter
}

// NewResponder initializes a new Responder with the http.ResponseWriter embedded.
func NewResponder(w http.ResponseWriter) *Responder {
	return &Responder{
		data:           nil,
		ResponseWriter: w,
	}
}

// WithData attaches data to the Responder which will be marshaled to JSON when the response is
// written.
func (r *Responder) WithData(data interface{}) {
	r.data = data
}

// WithHeader attaches a header to the Responder.
func (r *Responder) WithHeader(key, value string) {
	r.Header().Set(key, value)
}

func (r *Responder) write(statusCode int) {
	if r.data == nil {
		r.data = &GenericError{
			StatusCode: statusCode,
			StatusText: http.StatusText(statusCode),
		}
	}

	b, err := json.Marshal(r.data)
	if err != nil {
		r.WithData(nil)
		r.InternalServerError()
		return
	}

	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(statusCode)
	if _, err := r.Write(b); err != nil {
		r.InternalServerError()
	}
}
