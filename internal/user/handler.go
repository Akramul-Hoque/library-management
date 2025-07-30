package user

import (
	"encoding/json"
	"errors"
	"library-management/internal/middleware"
	"library-management/internal/response"
	"net/http"
	"strings"
)

var service = NewService()

// RegisterUserHandler handles user registration.
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload", "INVALID_PAYLOAD", nil)
		return
	}

	err := service.RegisterUser(&u)
	if err != nil {
		code := http.StatusInternalServerError
		msg := "Failed to register user: " + err.Error()
		msgCode := "REGISTER_ERROR"

		if err.Error() == "missing required fields" {
			code = http.StatusBadRequest
			msg = "Missing required fields: name, contact, password"
			msgCode = "MISSING_FIELDS"
		}

		response.Universal(w, code, false, msg, msgCode, nil)
		return
	}

	response.Universal(w, http.StatusCreated, true, "User registered successfully", "USER_REGISTERED", nil)
}

// OTPVerifyHandler handles OTP verification and login.
func OTPVerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Contact string `json:"contact"`
		OTP     string `json:"otp"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload", "INVALID_PAYLOAD", nil)
		return
	}

	req.Contact = strings.TrimSpace(req.Contact)
	req.OTP = strings.TrimSpace(req.OTP)

	err := service.VerifyOTP(req.Contact, req.OTP)
	if err != nil {
		response.Universal(w, http.StatusUnauthorized, false, err.Error(), "OTP_VERIFICATION_FAILED", nil)
		return
	}

	// user, err := service.AuthenticateUser(req.Contact, "")
	// if err != nil {
	// 	response.Universal(w, http.StatusUnauthorized, false, "User authentication failed after OTP verification", "AUTH_FAILED", nil)
	// 	return
	// }

	// token, err := middleware.GenerateJWT(user.ID, user.Role)
	// if err != nil {
	// 	response.Universal(w, http.StatusInternalServerError, false, "Failed to generate token", "JWT_ERROR", nil)
	// 	return
	// }

	response.Universal(w, http.StatusOK, true, "USER ACTIVATED SUCCEFULLY", "OTP_VERIFIED", nil)
}

// GetUsersHandler returns a list of users (without exposing passwords).
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	safeUsers := make([]UserResponse, 0, len(users))

	for _, u := range users {
		safeUsers = append(safeUsers, UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Contact:   u.Contact,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			IsActive:  u.IsActive,
		})
	}

	response.Universal(w, http.StatusOK, true, "Users retrieved", "USERS_RETRIEVED", safeUsers)
}

// SearchUsersHandler filters users by query parameters.
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

// LoginHandler handles user authentication and JWT generation.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload", "INVALID_PAYLOAD", nil)
		return
	}

	req.Contact = strings.TrimSpace(req.Contact)
	req.Password = strings.TrimSpace(req.Password)

	user, err := service.AuthenticateUser(req.Contact, req.Password)
	if err != nil {
		var inactiveErr ErrInactiveAccount
		if errors.As(err, &inactiveErr) {
			response.Universal(w, http.StatusForbidden, false, inactiveErr.Message, "ACCOUNT_INACTIVE_OTP_SENT", nil)
			return
		}

		response.Universal(w, http.StatusUnauthorized, false, "Invalid credentials", "INVALID_CREDENTIALS", nil)
		return
	}

	token, err := middleware.GenerateJWT(user.ID, user.Role)
	if err != nil {
		response.Universal(w, http.StatusInternalServerError, false, "Failed to generate token", "JWT_ERROR", nil)
		return
	}

	response.Universal(w, http.StatusOK, true, "Login successful", "LOGIN_SUCCESS", map[string]interface{}{
		"token": token,
	})
}
