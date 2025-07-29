package router

import (
	"library-management/internal/user"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	// @Summary Register a new user
	// @Description Add a new library user
	// @Tags meusermbers
	// @Accept json
	// @Produce json
	// @Param user body user.MemberRequest true "Member details"
	// @Success 201 {object} member.Member
	// @Router /api/members [post]
	r.HandleFunc("/api/user", user.RegisterUserHandler).Methods("POST")

	// @Summary List all members
	// @Description Retrieve all library members
	// @Tags members
	// @Produce json
	// @Success 200 {array} member.Member
	// @Router /api/members [get]
	r.HandleFunc("/api/user", user.GetUsersHandler).Methods("GET")

	// @Summary List all user by userType
	// @Description Retrieve all library user
	// @Tags users
	// @Produce json
	// @Success 200 {array} user.USERS
	// @Router /api/user [get]
	r.HandleFunc("/api/user/get-user", user.SearchUsersHandler).Methods("GET")
}
