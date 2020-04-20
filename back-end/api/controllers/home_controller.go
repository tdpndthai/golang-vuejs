package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/tdpndthai/golang-vuejs/api/auth"
	"github.com/tdpndthai/golang-vuejs/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	fmt.Println(tokenID)
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
