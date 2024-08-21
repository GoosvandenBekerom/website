package web

import (
	"net/http"
	"time"

	"github.com/goosvandenbekerom/website/data"
	"github.com/goosvandenbekerom/website/web/internal/assets"
	"github.com/goosvandenbekerom/website/web/internal/handlers"
)

const (
	serverTimeout = 10 * time.Second
)

func NewServer(storage *data.Storage) http.Handler {
	mux := http.NewServeMux()

	assets.Mount(mux)

	mux.Handle("GET /", handlers.Home(storage))

	return http.TimeoutHandler(mux, serverTimeout, "request timed out")
}
