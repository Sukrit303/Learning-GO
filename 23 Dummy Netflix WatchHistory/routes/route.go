package routes

import (
	"BackendNetflixWatchHistory/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/movies", controller.GetAllMovies).Methods("GET")

	router.HandleFunc("/api/v1/movies", controller.CreatingMovie).Methods("POST")

	router.HandleFunc("/api/v1/movies",controller.MarkAsWatched).Methods("POST")

	router.HandleFunc("/api/v1/movies",controller.DeleteAMovie).Methods("DELETE")

	return router
}