package cli

import (
	"fmt"

	"github.com/khaledez/txkv/kv"
)

type GetCommandHandler struct {
	key string
}

func (h *GetCommandHandler) ParseArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("GET command takes one argument")
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

type CountCommandHandler struct {
	value string
}

func (h *CountCommandHandler) ParseArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("COUNT command takes one argument")
	}
	h.value = args[0]
	return nil
}

func (h *CountCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	return fmt.Sprintf("%d", kvStore.Count(h.value))
}
