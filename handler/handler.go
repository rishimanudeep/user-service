package handler

import (
	"encoding/json"
	"net/http"
	"user-service/auth"
	"user-service/errors"
	"user-service/models"
)

type handler struct {
	userService UserService
}

func New(userService UserService) handler {
	return handler{userService: userService}
}

func (h handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUser(&user)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.userService.LoginUser(credentials.Email, credentials.Password)
	if err != nil {
		h.handleError(w, err)
		return
	}

	// Generate JWT token for the registered user
	token, err := auth.GenerateJWTToken(user.ID, user.Latitude, user.Longitude)
	if err != nil {
		http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
		return
	}

	// Create a response struct containing both the user data and the token
	response := struct {
		User  *models.User `json:"user"`
		Token string       `json:"token"`
	}{
		User:  user,
		Token: token,
	}

	// Return token in response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserProfile(email)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *handler) handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *errors.EntityNotFound:
		http.Error(w, e.Error(), http.StatusNotFound)
	case *errors.NoResponse:
		http.Error(w, e.Error(), http.StatusNotFound)
	case *errors.MissingParam:
		http.Error(w, e.Error(), http.StatusBadRequest)
	case *errors.ValidationError:
		http.Error(w, e.Error(), http.StatusBadRequest)
	case *errors.InternalServerError:
		http.Error(w, e.Error(), http.StatusInternalServerError)
	default:
		http.Error(w, "unknown error", http.StatusInternalServerError)
	}
}
