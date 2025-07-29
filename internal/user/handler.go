package user

import (
	"encoding/json"
	"library-management/response"
	"net/http"
)

var service = NewService()

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload", "INVALID_PAYLOAD", nil)
		return
	}

	err := service.RegisterUser(&u)
	if err != nil {
		if err.Error() == "missing required fields" {
			response.Universal(w, http.StatusBadRequest, false, "Missing required fields: name, contact, password", "MISSING_FIELDS", nil)
			return
		}
		response.Universal(w, http.StatusInternalServerError, false, "Failed to register user: "+err.Error(), "REGISTER_ERROR", nil)
		return
	}

	response.Universal(w, http.StatusCreated, true, "User registered successfully", "USER_REGISTERED", u)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	response.Universal(w, http.StatusOK, true, "Users retrieved", "USERS_RETRIEVED", users)
}

func SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	contact := r.URL.Query().Get("contact")
	email := r.URL.Query().Get("email")
	name := r.URL.Query().Get("name")
	userType := r.URL.Query().Get("userType")
	users := service.SearchUsers(contact, email, name, userType)
	if len(users) == 0 {
		response.Universal(w, http.StatusNotFound, false, "No users found for the given criteria", "NO_USERS_FOUND", nil)
		return
	}
	response.Universal(w, http.StatusOK, true, "Users retrieved", "USERS_RETRIEVED", users)
}
