package api

import (
    "github.com/gorilla/mux"
)

// SetupRouter sets up the routes for the application
func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/api/users", GetUserHandler).Methods("GET")
    r.HandleFunc("/api/some-route", SomeRouteHandler).Methods("GET")
    // Add other routes here
    return r
}
