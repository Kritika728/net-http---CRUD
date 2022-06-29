package handler

import (
	"encoding/json"
	"net/http"
	"nethttpcrud/db"
	"nethttpcrud/lib"
	"nethttpcrud/model"
	"nethttpcrud/pkg/utils"
)

// Get all movies

// response and request handlers
func GetMovies(w http.ResponseWriter, r *http.Request) {

	var movies []model.Movie

	db := db.SetupDB()
	lib.PrintMessage("Getting movies...")
	// Get all movies from movies table
	movies = utils.GetMoviesDetail(db)

	lib.EncoderResponse(w, movies)

}

//AddMovie : create
func AddMovie(w http.ResponseWriter, r *http.Request) {
	var response model.JsonResponse
	var movieRequest model.Movie

	if movieRequest, response = lib.DecoderRequest(r); response.Type != "Error" {
		response = utils.ValidateMovieData(movieRequest.MovieID, movieRequest.MovieName)
	}

	if response.Type != "Error" {

		utils.InsertMovieData(movieRequest)
		response = model.JsonResponse{Type: "success", Message: "Added successfully"}

	}
	json.NewEncoder(w).Encode(response)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movieRequest model.Movie
	var response model.JsonResponse

	movieRequest, response = lib.DecoderRequest(r)

	if response.Type != "Error" {
		Id := utils.GetMovieID(r)

		utils.UpdateMovie(movieRequest, Id)

		response = model.JsonResponse{Type: "Success", Message: "Successfully updated"}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	var response model.JsonResponse

	Id := utils.GetMovieID(r)

	utils.DeleteMovie(Id)

	response = model.JsonResponse{Type: "Success", Message: "Successfully deleted"}
	json.NewEncoder(w).Encode(response)
}
