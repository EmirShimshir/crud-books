package rest

import (
	"log"
	"net/http"
	"time"
)

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: [%s] - %s ", time.Now().Format(time.RFC822), r.Method, r.URL)
		next(w, r)
	}
}
