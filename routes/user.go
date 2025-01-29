package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"notes-api-go/db/database"

	"golang.org/x/crypto/bcrypt"
)

type UserRoutes struct {
	Queries *database.Queries
}

func (ur *UserRoutes) CreateUser(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	var user database.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Default role is "user" if not provided
	// if user.Roles == "" {
	// 	user.Roles = "user"
	// }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := ur.Queries.CreateUser(context.Background(), user); err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ur *UserRoutes) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := ur.Queries.GetUsers(context.Background())
	if err != nil {
		http.Error(w, "Could not fetch users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (ur *UserRoutes) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updateRequest struct {
		ID       int64    `json:"id"`
		Username string   `json:"username"`
		Password string   `json:"password"`
		Roles    []string `json:"roles"`
		Active   bool     `json:"active"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}
	updateRequest.Password = string(hashedPassword)

	if err := ur.Queries.UpdateUser(context.Background(), database.UpdateUserParams{
		ID:       updateRequest.ID,
		Username: updateRequest.Username,
		Password: updateRequest.Password,
		Roles:    updateRequest.Roles,
		Active:   updateRequest.Active,
	}); err != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ur *UserRoutes) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var deleteRequest struct {
		ID int64 `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&deleteRequest); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := ur.Queries.DeleteUser(context.Background(), deleteRequest.ID); err != nil {
		http.Error(w, "Could not delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
