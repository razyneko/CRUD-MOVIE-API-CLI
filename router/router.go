package router

import (
    "github.com/gorilla/mux"
	"github.com/movie-api/handlers"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
    r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
    r.HandleFunc("/movies", handlers.AddMovie).Methods("POST")
    r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
    r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")
    return r
}
