package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "postgres"
)

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Main function
func main() {

	// Get all movies
	http.HandleFunc("/get/movies/", GetMovies)

	//create a movie
	http.HandleFunc("/add/movies/", AddMovie)

	//update a movie
	http.HandleFunc("/update/movies/", UpdateMovie)

	//Delete a specific movie by the movieID
	http.HandleFunc("/delete/movies/", DeleteMovie)

	// serve the app
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Get all movies

// response and request handlers
func GetMovies(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting movies...")

	// Get all movies from movies table
	rows, err := db.Query("SELECT * FROM movies")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var movies []Movie

	// Foreach movie
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		checkErr(err)

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = JsonResponse{Type: "success", Data: movies}

	json.NewEncoder(w).Encode(response)
}

//create movie
func AddMovie(w http.ResponseWriter, r *http.Request) {
	//movieID := r.FormValue("movieid")
	//movieName := r.FormValue("moviename")
	var response JsonResponse
	var movieRequest Movie
	if err := json.NewDecoder(r.Body).Decode(&movieRequest); err != nil {
		response = JsonResponse{Type: "Error", Message: "Error in Parsing"}

	}

	if movieRequest.MovieID == "" || movieRequest.MovieName == "" {
		response = JsonResponse{Type: "Error", Message: "Value cann't be Blank"}
	} else {
		db := setupDB()
		var lastInsertID int
		printMessage("inserting the row")
		err := db.QueryRow("Insert into movies(movieid,moviename) values($1,$2) returning id", movieRequest.MovieID, movieRequest.MovieName).Scan(&lastInsertID)
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "added successfully"}

	}
	json.NewEncoder(w).Encode(response)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movieRequest Movie
	var response JsonResponse
	err := json.NewDecoder(r.Body).Decode(&movieRequest)

	checkErr(err)

	query := r.URL.Query()
	ID := query.Get("id")

	checkErr(err)
	//response=JsonResponse{Type:"Error",}

	db := setupDB()
	_, err = db.Exec("update movies set movieid =$1,moviename=$2 where id =$3 ", movieRequest.MovieID, movieRequest.MovieName, ID)
	checkErr(err)

	response = JsonResponse{Type: "Success", Message: "Successfully updated"}
	json.NewEncoder(w).Encode(response)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	var response JsonResponse

	query := r.URL.Query()
	ID := query.Get("id")

	db := setupDB()
	_, err := db.Exec("Delete from movies where id =$1 ", ID)
	checkErr(err)

	response = JsonResponse{Type: "Success", Message: "Successfully deleted"}
	json.NewEncoder(w).Encode(response)
}
