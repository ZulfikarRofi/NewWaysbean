package routes

import (
	"_project_name_/handlers"
	"_project_name_/pkg/middleware"
	"_project_name_/pkg/mysql"
	"_project_name_/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.Register)).Methods("POST")
}