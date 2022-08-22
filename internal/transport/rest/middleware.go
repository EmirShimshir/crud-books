package rest

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method": r.Method,
			"url":    r.URL,
		}).Info()
		next.ServeHTTP(w, r)
	})
}
