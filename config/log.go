package config

import (
	"errors"
	"log/slog"
	"os"
)

// setupSLog sets the default slog logger.
func setupSLog(cfg *Config) error {
	if cfg.Log == nil {
		return nil
	}
	opts := &slog.HandlerOptions{
		Level: cfg.Log.SLogLevel(),
	}
	if cfg.Debug {
		opts.Level = slog.LevelDebug
	}
	writer := os.Stdout
	var err error
	if cfg.Log.Output != "" && cfg.Log.Output != "stdout" {
		writer, err = os.OpenFile(cfg.Log.Output, os.O_APPEND, 0640)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
			writer, err = os.Create(cfg.Log.Output)
			if err != nil {
				return err
			}
		}
	}
	if cfg.Log.Bare && !cfg.Debug {
		opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey, slog.LevelKey, slog.SourceKey:
				return slog.Attr{Key: ""}
			case slog.MessageKey:
				return a
			}
			return a
		}
	} else if cfg.Log.Format != "json" && (cfg.Log.Output == "" || cfg.Log.Output == "stdout") {
		// Set default slog text handler
		return nil
	}

	var handler slog.Handler = slog.NewTextHandler(writer, opts)
	if cfg.Log.Format == "json" {
		handler = slog.NewJSONHandler(writer, opts)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return nil
}
