package router

import (
	"library-management/internal/member"

	"github.com/gorilla/mux"
)

func RegisterMemberRoutes(r *mux.Router) {
	// @Summary Register a new member
	// @Description Add a new library member
	// @Tags members
	// @Accept json
	// @Produce json
	// @Param member body member.MemberRequest true "Member details"
	// @Success 201 {object} member.Member
	// @Router /api/members [post]
	r.HandleFunc("/api/members", member.RegisterMemberHandler).Methods("POST")

	// @Summary List all members
	// @Description Retrieve all library members
	// @Tags members
	// @Produce json
	// @Success 200 {array} member.Member
	// @Router /api/members [get]
	r.HandleFunc("/api/members", member.GetMembersHandler).Methods("GET")
}
