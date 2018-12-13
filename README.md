[![travis ci status](https://travis-ci.org/kpurdon/apiresponse.svg?branch=master)](https://travis-ci.org/kpurdon/apiresponse)
[![codecov status](https://codecov.io/gh/kpurdon/apiresponse/branch/master/graph/badge.svg)](https://codecov.io/gh/kpurdon/apiresponse)
[![godoc](https://godoc.org/github.com/kpurdon/apiresponse?status.svg)](http://godoc.org/github.com/kpurdon/apiresponse)
[![Go Report Card](https://goreportcard.com/badge/github.com/kpurdon/apiresponse)](https://goreportcard.com/report/github.com/kpurdon/apiresponse)


apiresponse
-----

A simple helper library for writing JSON formatted responses in a HTTP  API.

## Examples

### Initialize a Responder and Write

``` go
func Foo(w http.ResponseWriter, r *http.Request) {
    responder := apiresponse.NewResponder(w)
    responder.OK()
}
```

### Initialize a Responder and Write with a Header

``` go
func Foo(w http.ResponseWriter, r *http.Request) {
    responder := apiresponse.NewResponder(w)
    responder.WithHeader("somekey", "somevalue")
    responder.OK()
}
```

### Initialize a Responder and Write with Data

``` go
func Foo(w http.ResponseWriter, r *http.Request) {
    responder := apiresponse.NewResponder(w)

    data := struct{Foo string `json:"foo"`}{Foo:"bar"}
    responder.WithData(data)

    responder.OK()
}
```
