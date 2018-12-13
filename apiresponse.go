//go:generate ./generate_helpers.sh

package apiresponse

import (
	"encoding/json"
	"net/http"
)

// Responder is the base type which all helper methods hang on. It should be initialized with the
// NewResponder method.
type Responder struct {
	w    http.ResponseWriter
	data interface{}
}

// NewResponder initializes a new Responder with the http.ResponseWriter embedded.
func NewResponder(w http.ResponseWriter) *Responder {
	return &Responder{
		w:    w,
		data: nil,
	}
}

// WithData attaches data to the Responder which will be marshaled to JSON when the response is
// written.
func (r Responder) WithData(data interface{}) {
	r.data = data
}

// WithHeader attaches a header to the Responder.
func (r Responder) WithHeader(key, value string) {
	r.w.Header().Set(key, value)
}

func (r Responder) write(statusCode int) {
	if r.data == nil {
		// TODO: something more useful?
		r.data = struct {
			StatusCode int    `json:"status_code"`
			StatusText string `json:"status_text"`
		}{
			StatusCode: statusCode,
			StatusText: http.StatusText(statusCode),
		}
	}

	b, err := json.Marshal(r.data)
	if err != nil {
		r.InternalServerError()
		return
	}

	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(statusCode)
	r.w.Write(b)
}
