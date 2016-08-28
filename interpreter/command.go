package interpreter

import (
	// "fmt"
)

func ExecGet(args []string) string {
	options := map[string]func([]string) string {
		"sql": GetSql,
	}

	// Make sure there are enough arguments provided.
	if len(args) < 1 {
		return "Not enough arguments."
	}

	// Check that the command exists in the options map, then execute it.
	if val, ok := options[args[0]]; ok {
		return val(args[1:])
	}

	return "That's not a valid target of the get command."
}

func GetSql(args []string) string {
	return `CREATE TABLE hamburgers`
}
