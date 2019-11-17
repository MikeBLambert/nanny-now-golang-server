package controllers

import (
	"net/http"

	"github.com/mikeblambert/nanny-now-golang-server/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
