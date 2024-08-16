package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	addr := flag.String("addr", "127.0.0.1:8080", "HTTP network address")
	flag.Parse()

	app := &application{
		logger: logger,
	}

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
