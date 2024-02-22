package main

import (
	"os"

	"log/slog"

	"pihole_exporter/config"
	"pihole_exporter/exporter"

	"pihole_exporter/server"
)

// Main application function
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	exporter, err := exporter.NewExporter(cfg)
	if err != nil {
		slog.Error("Error creating exporter", "error", err)
		os.Exit(1)
	}
	exporter.Register()

	if err := server.Run(cfg); err != nil {
		slog.Error("Error serving", "error", err)
		os.Exit(1)
	}
}
