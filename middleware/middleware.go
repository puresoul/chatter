package middleware

import (
	"fmt"
	"chatter/router"
	"net/http"
	"time"
)

func SetUp(h http.Handler) http.Handler {
	return router.ChainHandler(h, log, auth)
}

// Handler will log the HTTP requests.
func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        r.Header.Set("Access-Control-Allow-Origin", origin)
        r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        r.Header.Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
		h := r.Header.Get("Authorization")
		if h != "" {
		    fmt.Println(h)
		}
		next.ServeHTTP(w, r)
	})
}
