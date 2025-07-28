package member

import (
	"encoding/json"
	"net/http"
)

var service = NewService()

func RegisterMemberHandler(w http.ResponseWriter, r *http.Request) {
	var m Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := service.RegisterMember(&m)
	if err != nil {
		if err.Error() == "missing required fields" {
			http.Error(w, "Missing required fields: name, contact, password", http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to register member: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Member registered successfully"})
}

func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	members := service.GetAllMembers()
	json.NewEncoder(w).Encode(members)
}
