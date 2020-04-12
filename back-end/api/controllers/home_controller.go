package controllers

import (
	"net/http"

	"github.com/tdpndthai/golang-vuejs/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
