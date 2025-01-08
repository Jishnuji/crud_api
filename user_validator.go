package main

import (
	"context"
	"crud_it_krasava/storage"
	"encoding/json"
	"errors"
	"net/http"
)

type ContextKey string

const validatedUserKey ContextKey = "validatedUser"

func validateInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user storage.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger.Printf("Invalid input: %v", err)
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		if err := validateUser(user); err != nil {
			logger.Printf("Validation failed: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, validatedUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateUser(user storage.User) error {
	if user.Firstname == "" || user.Lastname == "" || user.Email == "" {
		return errors.New("missing required fields")
	}
	if user.Age <= 0 {
		return errors.New("invalid age")
	}
	return nil
}
