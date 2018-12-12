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
