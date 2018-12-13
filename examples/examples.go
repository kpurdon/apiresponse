package main

import (
	"log"
	"net/http"

	"github.com/kpurdon/apiresponse/v2"
)

func simple(w http.ResponseWriter, r *http.Request) {
	responder := apiresponse.NewResponder(w)
	responder.OK()

}

func withHeader(w http.ResponseWriter, r *http.Request) {
	responder := apiresponse.NewResponder(w)
	responder.WithHeader("Test-Key", "test/value")
	responder.OK()

}

func withData(w http.ResponseWriter, r *http.Request) {
	responder := apiresponse.NewResponder(w)
	responder.WithData(struct{ Name string }{Name: "test"})
	responder.OK()
}

func main() {
	http.HandleFunc("/simple", simple)
	http.HandleFunc("/with_header", withHeader)
	http.HandleFunc("/with_data", withData)
	log.Fatal(http.ListenAndServe(":3333", nil))
}
