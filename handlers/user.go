package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fixztter/rest-fundamentals/models"
	"github.com/fixztter/rest-fundamentals/repository"
	"github.com/fixztter/rest-fundamentals/server"
	"github.com/segmentio/ksuid"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHundler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Id:       id.String(),
			Email:    request.Email,
			Password: request.Password,
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(SignUpResponse{
			Id:    id.String(),
			Email: request.Email,
		})
	}
}
