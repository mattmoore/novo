package schema

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

type Schema struct {
	Databases []*Database `json:"databases"`
}

type Database struct {
	Connection sql.DB   `json:"-"`
	Name       string   `json:"name"`
	Tables     []*Table `json:"tables"`
}

type Table struct {
	Name        string    `json:"name"`
	PrimaryKeys []string  `json:"primary-keys"`
	Columns     []*Column `json:"columns"`
	Database    *Database `json:"-"`
}

type Column struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Table *Table `json:"-"`
}

func (s *Schema) Compile() {
	for _, db := range s.Databases {
		for _, table := range db.Tables {
			table.Database = db
			for _, c := range table.Columns {
				c.Table = table
			}
		}
	}
}

func (t *Table) GetPrimaryKeys() []*Column {
	keys := []*Column{}
	for _, c := range t.Columns {
		if c.IsPrimaryKey() {
			keys = append(keys, c)
		}
	}
	return keys
}

func (c *Column) IsPrimaryKey() bool {
	for _, key := range c.Table.PrimaryKeys {
		if c.Name == key {
			return true
		}
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
