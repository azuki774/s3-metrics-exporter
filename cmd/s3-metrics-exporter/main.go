package main

import (
	"log/slog"
	"os"
)

func main() {
	Execute()
}

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
