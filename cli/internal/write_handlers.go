package cli

import (
	"fmt"

	"github.com/khaledez/txkv/kv"
)

type SetCommandHandler struct {
	key, val string
}

func (h *SetCommandHandler) ParseArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("SET command takes two arguments")
	}
	h.key = args[0]
	h.val = args[1]
	return nil
}

func (h *SetCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	kvStore.Set(h.key, h.val)
	return ""
}

type DeleteCommandHandler struct {
	key string
}

func (h *DeleteCommandHandler) ParseArgs(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("DELETE command takes one argument")
	}
	h.key = args[0]
	return nil
}

func (h *DeleteCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	kvStore.Delete(h.key)
	return ""
}
