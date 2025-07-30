package response

import (
	"encoding/json"
	"net/http"
)

// APIResponse is a standard API response structure
type APIResponse struct {
	Status      string      `json:"status"`      // success / error
	Code        int         `json:"code"`        // HTTP status code
	Message     string      `json:"message"`     // Descriptive message
	MessageCode string      `json:"messageCode"` // Internal reference code
	Data        interface{} `json:"data"`        // Any result data
}

func Universal(w http.ResponseWriter, status int, success bool, message, messageCode string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := APIResponse{
		Status:      "success",
		Code:        status,
		Message:     message,
		MessageCode: messageCode,
		Data:        data,
	}
	json.NewEncoder(w).Encode(resp)
}

// The Universal function has been

// JSON writes a standard API response
func JSON(w http.ResponseWriter, code int, status, msg, msgCode string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := APIResponse{
		Status:      status,
		Code:        code,
		Message:     msg,
		MessageCode: msgCode,
		Data:        data,
	}

	json.NewEncoder(w).Encode(resp)
}
