package middleware

import (
	"net/http"
)

// CORSMiddleware allows CORS
type CORSMiddleware struct {
	Next http.Handler
}

// ServeHTTP is just CORS
func (cors CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	cors.Next.ServeHTTP(w, r)
}
