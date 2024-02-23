package test_data

import "log/slog"

var (
	LogLevel = []testData{
		{
			Name: "Empty_default_info",
			InputStr: []string{
				"",
			},
			ExpectedInt: []int{int(slog.LevelInfo)},
		},
		{
			Name: "Invalid_default_info",
			InputStr: []string{
				"invalid",
			},
			ExpectedInt: []int{int(slog.LevelInfo)},
		},
		{
			Name: "Info",
			InputStr: []string{
				"info",
			},
			ExpectedInt: []int{int(slog.LevelInfo)},
		},
		{
			Name: "Warn",
			InputStr: []string{
				"warn",
			},
			ExpectedInt: []int{int(slog.LevelWarn)},
		},
		{
			Name: "Error",
			InputStr: []string{
				"error",
			},
			ExpectedInt: []int{int(slog.LevelError)},
		},
		{
			Name: "Debug",
			InputStr: []string{
				"debug",
			},
			ExpectedInt: []int{int(slog.LevelDebug)},
		},
	}
)
