package test_data

import (
	"io"
	"log/slog"
)

type testData struct {
	Name         string
	InputStr     []string
	InputByte    [][]byte
	InputBool    []bool
	InputInt     []int
	ExpectedStr  []string
	ExpectedByte [][]byte
	ExpectedBool []bool
	ExpectedErr  []string
	ExpectedInt  []int
}

func SLogDisableOutput() {
	handler := slog.NewTextHandler(io.Discard, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
