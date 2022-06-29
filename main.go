package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kritika728/pkg/handler"
	_ "github.com/lib/pq"
)

// Main function
func main() {

	// Get all movies
	http.HandleFunc("/get/movies/", handler.GetMovies)

	//create a movie
	http.HandleFunc("/add/movies/", handler.AddMovie)

	//update a movie
	http.HandleFunc("/update/movies/", handler.UpdateMovie)

	//Delete a specific movie by the movieID
	http.HandleFunc("/delete/movies/", handler.DeleteMovie)

	// serve the app
	fmt.Println("Server at 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
