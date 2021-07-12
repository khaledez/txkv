package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/khaledez/txkv/kv"
)

func Cli(input io.Reader, output io.Writer, kvStore kv.KeyValueStore, prompt string) error {
	scanner := bufio.NewScanner(input)

	fmt.Fprint(output, prompt)
	for scanner.Scan() {
		instruction, err := parseCommandLine(scanner.Text())
		if err != nil {
			fmt.Fprintln(output, err)
			fmt.Fprint(output, prompt)
			continue
		}
		fmt.Fprintln(output, handleCommand(instruction))
		fmt.Fprint(output, prompt)
	}
	return nil
}

type Command string

const (
	GetCommand      = "GET"
	SetCommand      = "SET"
	DeleteCommand   = "DELETE"
	BeginCommand    = "BEGIN"
	CommitCommand   = "COMMIT"
	RollbackCommand = "ROLLBACK"
	CountCommand    = "COUNT"
)

type Instruction struct {
	Cmd  Command
	Args []string
}

func parseCommandLine(input string) (Instruction, error) {
	instruction := strings.Fields(input)

	switch strings.ToUpper(instruction[0]) {
	case GetCommand:
		if len(instruction[1:]) != 1 {
			return Instruction{}, fmt.Errorf("GET instruction takes one argument")
		}
		return Instruction{GetCommand, instruction[1:]}, nil
	case SetCommand:
		if len(instruction[1:]) != 2 {
			return Instruction{}, fmt.Errorf("SET instruction takes two arguments")
		}
		return Instruction{SetCommand, instruction[1:]}, nil
	case DeleteCommand:
		if len(instruction[1:]) != 1 {
			return Instruction{}, fmt.Errorf("DELETE instruction takes one argument")
		}
		return Instruction{DeleteCommand, instruction[1:]}, nil
	case CountCommand:
		if len(instruction[1:]) != 1 {
			return Instruction{}, fmt.Errorf("COUNT instruction takes one argument")
		}
		return Instruction{CountCommand, instruction[1:]}, nil
	case BeginCommand:
		return Instruction{BeginCommand, []string{}}, nil
	case CommitCommand:
		return Instruction{CommitCommand, []string{}}, nil
	case RollbackCommand:
		return Instruction{RollbackCommand, []string{}}, nil
	default:
		return Instruction{}, fmt.Errorf("unkown instruction %s", instruction[0])
	}
}
