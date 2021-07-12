package cli

func handleCommand(command Instruction) string {
	return string(command.Cmd)
}
