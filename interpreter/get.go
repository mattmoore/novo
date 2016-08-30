package interpreter

func ExecGetSql(args []string) string {
	return `CREATE TABLE hamburgers`
}

func ExecGetJson(args []string) string {
	return `{ "table": { "name": "hamburgers" } }`
}
