package interpreter

import (
	"strings"
)

func Parse(command string) string {
	if len(command) == 0 {
		return "No command given"
	}

	args := strings.Split(command, " ")

	if len(args) == 0 {
		return "Not enough args given."
	}

  commands := map[string]func([]string) string {
		"get": ExecGet,
  }

	if command, ok := commands[args[0]]; ok {
		return command(args[1:])
	}

	return "Invalid command."
}
