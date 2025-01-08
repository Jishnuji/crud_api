package main

import (
	"crud_it_krasava/storage"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	userStorage storage.UserStorage
}

func NewHandler(userStorage storage.UserStorage) *Handler {
	return &Handler{userStorage: userStorage}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(validatedUserKey).(storage.User)
	createdUser, err := h.userStorage.CreateUser(user)
	if err != nil {
		logger.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	logger.Printf("User created successfully: %v", createdUser.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

// getUser handles getting a user by ID
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		logger.Printf("Invalid user ID: %v", id)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userStorage.GetUserByID(id)
	if err != nil {
		logger.Printf("User not found: %v", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	logger.Printf("User fetched successfully: %v", user.ID)
	json.NewEncoder(w).Encode(user)
}

// updateUser handles updating a user by ID
func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		logger.Printf("Invalid user ID: %v", idStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := r.Context().Value(validatedUserKey).(storage.User)
	updatedUser, err := h.userStorage.UpdateUser(id, user)
	if err != nil {
		logger.Printf("Failed to update user: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	logger.Printf("User updated successfully: %v", id)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(updatedUser)
}
