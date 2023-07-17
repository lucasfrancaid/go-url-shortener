package standard_middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		ResponseMiddleware(w, r)
	})
}

func ResponseMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v %v \n", w.Header().Get("Status"), r.Method, r.Host, r.URL)
}
