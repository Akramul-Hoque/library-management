package member

import (
	"encoding/json"
	"net/http"
)

func RegisterMemberHandler(w http.ResponseWriter, r *http.Request) {
	var m Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	RegisterMember(m.Name)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	members := GetAllMembers()
	json.NewEncoder(w).Encode(members)
}
