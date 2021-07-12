package cli

import (
	"fmt"

	"github.com/khaledez/txkv/kv"
)

type CommandHandler interface {
	ParseArgs(args []string) error
	Handle(kvStore kv.KeyValueStore) string
}

type GetCommandHandler struct {
	key string
}

func (h *GetCommandHandler) ParseArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("GET instruction takes one argument")
	}
	h.key = args[0]
	return nil
}

func (h *GetCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	if val, ok := kvStore.Get(h.key); ok {
		return val
	} else {
		return "key not set"
	}
}

// --- Transaction Commands handlers
type BeginCommandHandler struct{}

func (h BeginCommandHandler) ParseArgs(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("BEGIN command doesn't accept arguments")
	}

	return nil
}

func (h BeginCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	err := kvStore.BeginTransaction()
	if err != nil {
		return err.Error()
	}
	return ""
}

type CommitCommandHandler struct{}

func (h CommitCommandHandler) ParseArgs(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("BEGIN command doesn't accept arguments")
	}

	return nil
}

func (h CommitCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	err := kvStore.BeginTransaction()
	if err != nil {
		return err.Error()
	}
	return ""
}
