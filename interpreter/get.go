package interpreter

import (
	"github.com/mattmoore/novo/schema"
)

func ExecGetSql(args []string) string {
	return schema.Parse(args[0]).Sql()
}

func ExecGetJson(args []string) string {
	return schema.Parse(args[0]).Json()
}
