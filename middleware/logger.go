package middleware

import (
	"log"
	"net/http"

	"github.com/amnay-mo/kanban-api/utils"
)

// LoggerMiddleware is a logging middleware
type LoggerMiddleware struct {
	Next http.Handler
}

// Handle is LoggerMiddleware's implementation
func (l LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lrw := utils.NewLoggingResponseWriter(w)
	l.Next.ServeHTTP(lrw, r)
	log.Printf("%s %d %s %s %s", r.Method, lrw.StatusCode, r.URL, lrw.CurrentUser, r.Header.Get("X-Forwarded-For"))
}
