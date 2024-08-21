package main

import (
	"flag"
	"log/slog"
	"net/http"

	"github.com/goosvandenbekerom/website/data"
	"github.com/goosvandenbekerom/website/pkg/logger"
	"github.com/goosvandenbekerom/website/web"
)

var (
	addr = flag.String("addr", ":8080", "address to host webserver at")
)

func main() {
	flag.Parse()
	logger.Initialize()

	storage := data.NewStorage()
	server := web.NewServer(storage)

	slog.Info("starting http server", slog.String("addr", *addr))
	if err := http.ListenAndServe(*addr, server); err != nil {
		slog.Error("failed to start http server", slog.String("error", err.Error()))
	}
}
