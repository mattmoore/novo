package main

import "fmt"
import "os"
import "strings"
import "github.com/mattmoore/novo/interpreter"

import "github.com/mattmoore/novo/schema"

func main() {
	command := strings.Join(os.Args[1:], " ")
	response := interpreter.ExecCommand(command)
	fmt.Println(response)

	// db := schema.Database{Name: "db1"}
	// table1 := schema.Table{
	// 	Name:        "table1",
	// 	PrimaryKeys: []string{"id"},
	// 	Columns: []*schema.Column{
	// 		&schema.Column{Name: "id", Type: "serial"},
	// 		&schema.Column{Name: "col1", Type: "varchar(100)"},
	// 	},
	// }
	// db.Tables = append(db.Tables, &table1)
	// db.Compile()
	// fmt.Println(db.Sql())
}
