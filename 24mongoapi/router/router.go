package router

import (
	"github.com/divyamraj1301/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMoviesFromDB).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovieInDB).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatchedInDB).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovieFromDB).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMoviesFromDB).Methods("DELETE")

	return router
}
