package middleware

import (
	"fmt"
	"chatter/router"
	"net/http"
	"time"
)

func SetUp(h http.Handler) http.Handler {
	return router.ChainHandler(h, log)
}

// Handler will log the HTTP requests.
func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
