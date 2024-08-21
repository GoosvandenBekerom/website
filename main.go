package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

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

	storage, err := data.NewStorage()
	if err != nil {
		fail("init storage", err)
	}

	server := web.NewServer(storage)

	slog.Info("starting http server", slog.String("addr", *addr))
	if err := http.ListenAndServe(*addr, server); err != nil {
		fail("start http server", err)
	}
}

func fail(action string, err error) {
	slog.Error("failed to "+action, slog.String("error", err.Error()))
	os.Exit(1)
}
