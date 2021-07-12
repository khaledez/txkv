package main

import (
	"os"

	"github.com/khaledez/txkv/cli"
	"github.com/khaledez/txkv/kv"
)

func main() {
	cli.Cli(os.Stdin, os.Stdout, kv.NewMemoryStore(), "> ")
}
