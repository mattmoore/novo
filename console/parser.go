package console

import (
	"strings"
)

func Parse(command string) string {
	args := strings.Split(command, " ")

  commands := map[string]func([]string) string {
		"get": ExecGet,
  }

	return commands[args[0]](args[1:])
	return ""
}
