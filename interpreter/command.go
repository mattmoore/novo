package interpreter

import (
	"fmt"
	"strings"
)

func ExecCommand(command string) string {
	if len(command) == 0 {
		return "No command given"
	}

	args := strings.Split(command, " ")
	if len(args) == 0 {
		return "No args presented"
	}

	commands := map[string]func([]string) string{
		"dump": ExecDump,
		"get":  ExecGet,
	}

	if command, ok := commands[args[0]]; ok {
		return command(args[1:])
	}

	return "Invalid command."
}

func ExecDump(args []string) string {
	host := args[0]
	dbname := args[1]
	format := args[2]
	if len(format) > 0 {
		return DumpFormat(host, dbname, format)
	}
	return fmt.Sprintf("Invalid dump format [%s]", args[0])
}

func ExecGet(args []string) string {
	commands := map[string]func([]string) string{
		"sql":  ExecGetSql,
		"json": ExecGetJson,
	}

	if len(args) == 0 {
		return "No args presented for get command"
	}

	return CheckAndExecCommand(args, commands)
}

func CheckAndExecCommand(args []string, commands map[string]func([]string) string) string {
	if command, ok := commands[args[0]]; ok {
		return command(args[1:])
	}
	return "That's not a valid target of the get command."
}
