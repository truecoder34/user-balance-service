package controllers

import (
	"net/http"

	"github.com/truecoder34/user-balance-service/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To API")

}
