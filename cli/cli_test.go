package cli

import (
	"io"
	"strings"
	"testing"

	"github.com/khaledez/txkv/kv"
	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	testCases := []struct {
		TestTitle           string
		InputLines          []string
		ExpectedOutputLines []string
	}{
		{
			TestTitle:           "reading entry not exist",
			InputLines:          []string{"GET foo"},
			ExpectedOutputLines: []string{"key not set"},
		},
		{
			TestTitle:           "setting a value then reading it",
			InputLines:          []string{"SET foo 123", "GET foo"},
			ExpectedOutputLines: []string{"123"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestTitle, func(t *testing.T) {
			out := &mockWriter{}
			err := Cli(mockReader(testCase.InputLines), out, kv.NewMemoryStore(), "")

			assert.NoError(t, err)
			assert.Equal(t, testCase.ExpectedOutputLines, out.Lines()[:len(out.Lines())-1])
		})
	}
}

func mockReader(lines []string) io.Reader {
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
