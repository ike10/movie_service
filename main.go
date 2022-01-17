package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ike10/movie_service/config"
	"github.com/ike10/movie_service/dao"
	"github.com/ike10/movie_service/handlers"
)

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", handlers.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", handlers.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", handlers.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", handlers.FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatal(err)
	}
}
