package web

import (
	"net/http"
	"time"
)

const (
	serverTimeout = 10 * time.Second
)

func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Landing Page!"))
	})

	return http.TimeoutHandler(mux, serverTimeout, "request timed out")
}
