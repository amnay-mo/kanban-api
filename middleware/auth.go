package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/amnay-mo/kanban-api/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var secretKey []byte

// SetSecretKey sets the jwt token signing key
func SetSecretKey(key string) {
	secretKey = []byte(key)
}

// GetSecretKey returns the jwt token signing key
func GetSecretKey() []byte {
	return secretKey
}

func parseToken(token string) (*jwt.Token, error) {
	if token == "" {
		return nil, fmt.Errorf("token is empty")
	}
	keyFunc := func(_ *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}
	return jwt.Parse(token, keyFunc)
}

func isValid(token *jwt.Token) bool {
	return int64(token.Claims.(jwt.MapClaims)["exp"].(float64)) > time.Now().Unix()
}

func getEmail(token *jwt.Token) string {
	email := token.Claims.(jwt.MapClaims)["user"].(string)
	return email
}

func extractToken(authHeader string) string {
	splitParts := strings.Split(authHeader, " ")
	if len(splitParts) != 2 {
		return ""
	}
	if splitParts[0] != "Bearer" {
		return ""
	}
	return splitParts[1]
}

// AuthMiddleware is AuthMiddleware's implementation of the Handler interface
func AuthMiddleware(Next func(w http.ResponseWriter, r *http.Request, p httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		tokenString := extractToken(r.Header.Get("Authorization"))
		token, err := parseToken(tokenString)
		if err != nil {
			utils.Jsonify(w, r, map[string]string{"error": "Bad Token"}, http.StatusUnauthorized)
			return
		}
		if isValid(token) {
			if _, ok := w.(*utils.LoggingResponseWriter); ok {
				w.(*utils.LoggingResponseWriter).CurrentUser = getEmail(token)
			}
			Next(w, r, p)
		} else {
			utils.Jsonify(w, r, map[string]string{"error": "Token Expired"}, http.StatusUnauthorized)
		}
	}
}

// JWTMiddleware is the auth middleware
// var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })
