package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	//get port from environment
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe("0.0.0.0:"+port, nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uri: %s", req.URL)
		next.ServeHTTP(w, req)
	})
}
