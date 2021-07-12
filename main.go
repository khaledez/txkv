package main

import (
	"os"

	"github.com/khaledez/txkv/cli"
	"github.com/khaledez/txkv/kv"
)

func main() {
	// TODO: this is not a production ready main function, it doesn't handle signals, but that's not needed for the current scope
	cli.Cli(os.Stdin, os.Stdout, kv.NewMemoryStore(), "> ")
}
