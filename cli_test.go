package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reading_entry_not_exists(t *testing.T) {
	// Given ...
	inputLines := []string{"GET foo"}
	expectedOutput := "key not set"

	// When ...
	out := &mockWriter{}
	err := Cli(MockReader(inputLines), out)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, out.Lines()[0])
}

func MockReader(lines []string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}

type mockWriter struct {
	buffer string
}

func (w *mockWriter) Lines() []string {
	return strings.Split(w.buffer, "\n")
}

func (w *mockWriter) Write(p []byte) (int, error) {
	w.buffer += string(p)

	return len(p), nil
}
