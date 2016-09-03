package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Schema struct {
	Databases []*Database `json:"databases"`
}

type Database struct {
	Name   string   `json:"name"`
	Tables []*Table `json:"tables"`
}

type Table struct {
	Name       string    `json:"name"`
	PrimaryKey string    `json:"primary-key"`
	Columns    []*Column `json:"columns"`
}

type Column struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Table *Table `json:"-"`
}

func (t *Table) GetPrimaryKey() *Column {
	for _, c := range t.Columns {
		if c.Name == t.PrimaryKey {
			return c
		}
	}
	return nil
}

func (c *Column) IsPrimaryKey() bool {
	if c.Name == c.Table.PrimaryKey {
		return true
	}
	return false
}

// SQL export

func (s *Schema) Sql() string {
	var buffer bytes.Buffer
	for _, db := range s.Databases {
		var tables []string
		for _, table := range db.Tables {
			tables = append(tables, table.Sql())
		}
		buffer.WriteString(strings.Join(tables, "\n"))
	}
	return buffer.String()
}

func (t *Table) Sql() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", t.Name))
	columns := []string{}
	for _, c := range t.Columns {
		columns = append(columns, fmt.Sprintf("  %s", c.Sql()))
	}
	buffer.WriteString(strings.Join(columns, ",\n"))
	buffer.WriteString("\n);")
	return buffer.String()
}

func (c *Column) Sql() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s", c.Name, strings.ToUpper(c.Type)))
	if c.IsPrimaryKey() {
		buffer.WriteString(fmt.Sprintf(" PRIMARY KEY"))
	}
	return buffer.String()
}

// JSON export

func (s *Schema) Json() string {
	buffer, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return ""
	}
	return string(buffer)
}
