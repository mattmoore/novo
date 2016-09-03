package main

import "fmt"
import "os"
import "strings"
import "github.com/mattmoore/novo/interpreter"

// import "github.com/mattmoore/novo/schema"

func main() {
	command := strings.Join(os.Args[1:], " ")
	response := interpreter.ExecCommand(command)
	fmt.Println(response)

	// schema1 := new(schema.Schema)
	//
	// db1 := schema.Database{Name: "db1"}
	// schema1.Databases = append(schema1.Databases, &db1)
	//
	// table1 := schema.Table{
	// 	Name: "table1",
	// 	Columns: []*schema.Column{
	// 		&schema.Column{Name: "id", Type: "serial", PrimaryKey: true},
	// 		&schema.Column{Name: "col1", Type: "varchar(100)"},
	// 	},
	// }
	//
	// db1.Tables = append(db1.Tables, &table1)
	//
	// fmt.Print(schema1.GetSchemaSql())
}
