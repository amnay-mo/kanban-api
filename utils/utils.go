package utils

import (
	"encoding/json"
	"net/http"
)

// Jsonify does the trick!
func Jsonify(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

// LoggingResponseWriter is useful for getting logs
type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode  int
	CurrentUser string
}

// NewLoggingResponseWriter returns a pointer to a LoggingResponseWriter
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	lrw := new(LoggingResponseWriter)
	lrw.ResponseWriter = w
	return lrw
}

// WriteHeader writes the status code
func (lrw *LoggingResponseWriter) WriteHeader(status int) {
	lrw.ResponseWriter.WriteHeader(status)
	lrw.StatusCode = status
}

// SetCurrentUser is useful for logging the user's email
func (lrw *LoggingResponseWriter) SetCurrentUser(email string) {
	lrw.CurrentUser = email
}
