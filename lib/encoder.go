package lib

import (
	"encoding/json"
	"net/http"
	"nethttpcrud/model"
)

func EncoderResponse(w http.ResponseWriter, movies []model.Movie) {
	var response = model.JsonResponse{Type: "success", Data: movies}
	json.NewEncoder(w).Encode(response)
}

func DecoderRequest(r *http.Request) (movieRequest model.Movie, response model.JsonResponse) {
	if err := json.NewDecoder(r.Body).Decode(&movieRequest); err != nil {
		response = model.JsonResponse{Type: "Error", Message: "Error in Parsing"}

	}
	return
}
