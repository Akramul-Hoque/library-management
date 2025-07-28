package member

import (
	"encoding/json"
	"library-management/response"
	"net/http"
)

var service = NewService()

func RegisterMemberHandler(w http.ResponseWriter, r *http.Request) {
	var m Member
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		response.Universal(w, http.StatusBadRequest, false, "Invalid request payload")
		return
	}

	err := service.RegisterMember(&m)
	if err != nil {
		if err.Error() == "missing required fields" {
			response.Universal(w, http.StatusBadRequest, false, "Missing required fields: name, contact, password")
			return
		}
		response.Universal(w, http.StatusInternalServerError, false, "Failed to register member: "+err.Error())
		return
	}

	response.Universal(w, http.StatusCreated, true, "Member registered successfully")
}

func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	members := service.GetAllMembers()
	json.NewEncoder(w).Encode(members)
}
