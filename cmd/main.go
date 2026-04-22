package main

import (
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addrs: ":8080",
		db:    dbConfig{},
	}

	api := &application{
		config: cfg,
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	h := api.mount()

	if err := api.run(h); err != nil {
		slog.Error("Failed to start the server", "error", err)
		os.Exit(1)
	}
}
