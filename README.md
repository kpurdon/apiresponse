[![CircleCI](https://circleci.com/gh/kpurdon/apiresponse.svg?style=svg)](https://circleci.com/gh/kpurdon/apiresponse)
[![codecov status](https://codecov.io/gh/kpurdon/apiresponse/branch/master/graph/badge.svg)](https://codecov.io/gh/kpurdon/apiresponse)
[![godoc](https://godoc.org/github.com/kpurdon/apiresponse?status.svg)](http://godoc.org/github.com/kpurdon/apiresponse)
[![Go Report Card](https://goreportcard.com/badge/github.com/kpurdon/apiresponse)](https://goreportcard.com/report/github.com/kpurdon/apiresponse)


apiresponse
-----

A simple helper for writing JSON formatted responses in an HTTP API.

## Examples

### Initialize a Responder and Write

``` go
func Foo(w http.ResponseWriter, r *http.Request) {
    responder := apiresponse.NewResponder(w)
    responder.OK()
}
```

results in

``` json
HTTP/1.1 200 OK
Content-Length: 38
Content-Type: application/json
Date: Thu, 13 Dec 2018 18:56:14 GMT

{
    "status_code": 200,
    "status_text": "OK"
}
```

### Initialize a Responder and Write with a Header

``` go
func Foo(w http.ResponseWriter, r *http.Request) {
    responder := apiresponse.NewResponder(w)
    responder.WithHeader("Test-Key", "test/value")
    responder.OK()
}
```

results in

``` json
HTTP/1.1 200 OK
Content-Length: 38
Content-Type: application/json
Date: Thu, 13 Dec 2018 18:56:17 GMT
Test-Key: test/value

{
    "status_code": 200,
    "status_text": "OK"
}
```

### Initialize a Responder and Write with Data

``` go
func Foo(w http.ResponseWriter, r *http.Request) {
    responder := apiresponse.NewResponder(w)

    data := struct{Name string}{Name:"test"}
    responder.WithData(data)

    responder.OK()
}
```

results in

``` json
HTTP/1.1 200 OK
Content-Length: 15
Content-Type: application/json
Date: Thu, 13 Dec 2018 18:56:21 GMT

{
    "Name": "test"
}
```
