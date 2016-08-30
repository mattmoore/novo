package schema

import (
	"fmt"
	"strings"
)

type Schema struct {
	Databases []*Database
}

type Database struct {
	Name   string
	Tables []*Table
}

type Table struct {
	Name    string
	Columns []*Column
}

type Column struct {
	Name       string
	Type       string
	PrimaryKey bool
}

func (s *Schema) GetSchemaSql() string {
	var sql string
	for _, db := range s.Databases {
		for _, table := range db.Tables {
			sql += table.GetTableSql()
		}
	}
	return sql
}

func (t *Table) GetTableSql() string {
	var sql string
	sql += fmt.Sprintf("CREATE TABLE %s (\n", t.Name)
	columns := []string{}
	for _, c := range t.Columns {
		columns = append(columns, fmt.Sprintf("  %s", c.GetColumnSql()))
	}
	sql += strings.Join(columns, ",\n")
	sql += "\n);\n"
	return sql
}

func (c *Column) GetColumnSql() string {
	var sql string
	sql += fmt.Sprintf("%s %s", c.Name, strings.ToUpper(c.Type))
	if c.PrimaryKey {
		sql += fmt.Sprintf(" PRIMARY KEY")
	}
	return sql
}
