package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/amnay-mo/kanban-api/middleware"
	"github.com/amnay-mo/kanban-api/model"
	"github.com/amnay-mo/kanban-api/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

// SignUp is a handler for the signin route
func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, err := decodeUser(r)
	if err != nil {
		log.Printf("Could not decode user: %v", err)
		response := map[string]string{
			"error": "Bad request payload",
		}
		utils.Jsonify(w, r, response, http.StatusBadRequest)
	}
	err = model.Signup(user)
	log.Println(user)
}

// Authenticate is for auth
func Authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, _ := decodeUser(r)
	if !model.Authenticate(user) {
		utils.Jsonify(w, r, map[string]string{"error": "Bad Credentials"}, http.StatusUnauthorized)
	} else {
		GetToken(w, r, user)
	}
}

// GetToken returns a JWT token
func GetToken(w http.ResponseWriter, r *http.Request, user *model.User) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var mySigningKey = middleware.GetSecretKey()
	claims["admin"] = true
	claims["user"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, _ := token.SignedString(mySigningKey)
	log.Printf("[GetToken] %v", tokenString)
	utils.Jsonify(w, r, map[string]string{"token": tokenString}, http.StatusOK)
}
