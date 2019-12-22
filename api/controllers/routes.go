package controllers

import "github.com/mikeblambert/nanny-now-golang-server/api/middlewares"

func (server *Server) initializeRoutes() {

	// Home Route
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	// Login Route
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	//Agency routes
	server.Router.HandleFunc("/agencies", middlewares.SetMiddlewareJSON(server.CreateAgency)).Methods("POST")
	server.Router.HandleFunc("/agencies", middlewares.SetMiddlewareJSON(server.GetAgencies)).Methods("GET")
	server.Router.HandleFunc("/agencies/{id}", middlewares.SetMiddlewareJSON(server.GetAgency)).Methods("GET")
	server.Router.HandleFunc("/agencies/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateAgency))).Methods("PUT")
	server.Router.HandleFunc("/agencies/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteAgency)).Methods("DELETE")

	//Users routes
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.DeleteUser)).Methods("DELETE")
}
