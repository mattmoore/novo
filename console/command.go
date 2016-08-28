package console

import (
	// "fmt"
)

func ExecGet(args []string) string {
	options := map[string]func([]string) string {
		"sql": GetSql,
	}
	return options[args[0]](args[1:])
	return ""
}

func GetSql(args []string) string {
	return `CREATE TABLE
hamburgers`
}
