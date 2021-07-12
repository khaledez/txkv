package cli

import (
	"fmt"

	"github.com/khaledez/txkv/kv"
)

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
		return fmt.Errorf("COMMIT command doesn't accept arguments")
	}

	return nil
}

func (h CommitCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	err := kvStore.CommitTransaction()
	if err != nil {
		return err.Error()
	}
	return ""
}

type RollbackCommandHandler struct{}

func (h RollbackCommandHandler) ParseArgs(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("ROLLBACK command doesn't accept arguments")
	}

	return nil
}

func (h RollbackCommandHandler) Handle(kvStore kv.KeyValueStore) string {
	err := kvStore.RollbackTransaction()
	if err != nil {
		return err.Error()
	}
	return ""
}
