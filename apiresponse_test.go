package apiresponse

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewResponder(t *testing.T) {
	actual := NewResponder(httptest.NewRecorder())
	require.NotNil(t, actual)
	assert.Nil(t, actual.data)
}

func TestResponderWithData(t *testing.T) {
	actual := NewResponder(httptest.NewRecorder())
	actual.WithData("test")
	assert.Equal(t, "test", actual.data)
}

func TestResponderWithHeader(t *testing.T) {
	actual := NewResponder(httptest.NewRecorder())
	actual.WithHeader("testkey", "testvalue")
	assert.Equal(t, "testvalue", actual.Header().Get("testkey"))
}

func TestResponderWrite(t *testing.T) {
	testCases := []struct {
		label          string
		modFunc        func(r *Responder)
		expectedCode   string
		expectedHeader http.Header
		expectedBody   string
	}{
		{
			label:   "no modifiers",
			modFunc: func(r *Responder) {},
			expectedHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			expectedBody: `{"status_code":200,"status_text":"OK"}`,
		}, {
			label: "with data",
			modFunc: func(r *Responder) {
				r.WithData(struct{ Test string }{Test: "test"})
			},
			expectedHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			expectedBody: `{"Test":"test"}`,
		}, {
			label: "with header",
			modFunc: func(r *Responder) {
				r.WithHeader("testkey", "testvalue")
			},
			expectedHeader: http.Header{
				"Content-Type": []string{"application/json"},
				"testkey":      []string{"testvalue"},
			},
			expectedBody: `{"status_code":200,"status_text":"OK"}`,
		}, {
			label: "with header and data",
			modFunc: func(r *Responder) {
				r.WithData(struct{ Test string }{Test: "test"})
				r.WithHeader("testkey", "testvalue")
			},
			expectedHeader: http.Header{
				"Content-Type": []string{"application/json"},
				"testkey":      []string{"testvalue"},
			},
			expectedBody: `{"Test":"test"}`,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Log(tc.label)
			t.Parallel()

			w := httptest.NewRecorder()
			actual := NewResponder(w)

			tc.modFunc(actual)
			actual.OK()

			assert.Equal(t, 200, w.Code)
			for k, v := range tc.expectedHeader {
				assert.Equal(t, v[0], w.Header().Get(k))
			}
			assert.Equal(t, tc.expectedBody, w.Body.String())
		})
	}
}
