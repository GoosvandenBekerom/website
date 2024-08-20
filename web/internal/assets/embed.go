package assets

import (
	"embed"
	"net/http"
)

//go:embed all:dist
var Assets embed.FS

// Mount the embedded assets on the given serve mux
func Mount(mux *http.ServeMux) {
	mux.Handle("GET /dist/", http.FileServer(http.FS(Assets)))
}
