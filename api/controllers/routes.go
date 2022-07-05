package controllers

import "github.com/truecoder34/user-balance-service/api/middlewares"

func (s *Server) initializeRoutes() {
	//base login
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// user methods
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")

	// account methods
	s.Router.HandleFunc("/accounts", middlewares.SetMiddlewareJSON(s.CreateAccount)).Methods("POST")
	s.Router.HandleFunc("/accounts", middlewares.SetMiddlewareJSON(s.GetAccounts)).Methods("GET")
	s.Router.HandleFunc("/accounts/{id}", middlewares.SetMiddlewareJSON(s.GetAccount)).Methods("GET")

	// money transfers
	s.Router.HandleFunc("/money", middlewares.SetMiddlewareJSON(s.AddRemoveMoney)).Methods("POST")
}
