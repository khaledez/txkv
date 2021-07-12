package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	cli "github.com/khaledez/txkv/cli/internal"
	"github.com/khaledez/txkv/kv"
)

func Cli(input io.Reader, output io.Writer, kvStore kv.KeyValueStore, prompt string) error {
	scanner := bufio.NewScanner(input)

	fmt.Fprint(output, prompt)
	for scanner.Scan() {
		handler, err := parseCommandLine(scanner.Text())
		if err != nil {
			fmt.Fprintln(output, err)
			fmt.Fprint(output, prompt)
			continue
		}
		result := handler.Handle(kvStore)
		if len(result) > 0 {
			fmt.Fprintln(output, result)
		}
		fmt.Fprint(output, prompt)
	}
	return nil
}

type Command string

type CommandHandler interface {
	ParseArgs(args []string) error
	Handle(kvStore kv.KeyValueStore) string
}

const (
	GetCommand      = "GET"
	SetCommand      = "SET"
	DeleteCommand   = "DELETE"
	BeginCommand    = "BEGIN"
	CommitCommand   = "COMMIT"
	RollbackCommand = "ROLLBACK"
	CountCommand    = "COUNT"
)

var commands = map[Command]CommandHandler{
	GetCommand:      &cli.GetCommandHandler{},
	SetCommand:      &cli.SetCommandHandler{},
	DeleteCommand:   &cli.DeleteCommandHandler{},
	CountCommand:    &cli.CountCommandHandler{},
	BeginCommand:    &cli.BeginCommandHandler{},
	CommitCommand:   &cli.CommitCommandHandler{},
	RollbackCommand: &cli.RollbackCommandHandler{},
}

func parseCommandLine(input string) (CommandHandler, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("")
	}
	instruction := strings.Fields(input)

	commandStr := strings.ToUpper(instruction[0])

	if handler, ok := commands[Command(commandStr)]; ok {
		return handler, handler.ParseArgs(instruction[1:])
	}
	return nil, fmt.Errorf("unknown command %s", commandStr)
}
