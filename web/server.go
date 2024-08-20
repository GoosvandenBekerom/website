package web

import "net/http"

func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Landing Page!"))
	})

	return mux
}
