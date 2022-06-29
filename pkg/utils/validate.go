package utils

import (
	"net/http"
	"nethttpcrud/model"
)

func ValidateMovieData(id, name string) (response model.JsonResponse) {

	if id == "" || name == "" {
		response = model.JsonResponse{Type: "Error", Message: "Value cann't be Blank"}
	}
	return
}

func GetMovieID(r *http.Request) (Id string) {
	query := r.URL.Query()
	Id = query.Get("id")
	return
}
