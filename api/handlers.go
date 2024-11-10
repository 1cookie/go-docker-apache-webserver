package api

import (
    "fmt"
    "net/http"
)

// Handler function to respond to a request
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // Just an example response
    fmt.Fprintf(w, "User data fetched!")
}
func SomeRouteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the response for /api/some-route")
}