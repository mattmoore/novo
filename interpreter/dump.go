package interpreter

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mattmoore/novo/schema"
)

func Connect(host string, dbname string) (sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s dbname=%s sslmode=disable", host, dbname))
	if err != nil {
		fmt.Println(err)
	}
	return *db, err
}

func Dump(host string, dbname string) *schema.Schema {
	conn, _ := Connect(host, dbname)
	db := &schema.Database{
		Connection: conn,
		Name:       dbname,
	}
	GetTables(db)
	conn.Close()
	s := &schema.Schema{
		Databases: []*schema.Database{db},
	}
	s.Compile()
	return s
}

func DumpFormat(host string, dbname string, format string) string {
	s := Dump(host, dbname)
	switch format {
	case "json":
		return s.Json()
	case "sql":
		return s.Sql()
	default:
		return s.Json()
	}
}

func GetTables(db *schema.Database) {
	rows, err := db.Connection.Query("SELECT tablename FROM pg_tables WHERE tableowner = $1", db.Name)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		table := &schema.Table{
			Name: tableName,
		}
		db.Tables = append(db.Tables, table)
	}

	for _, table := range db.Tables {
		GetColumns(db, table)
		for _, key := range GetPrimaryKeys(db, table.Name) {
			table.PrimaryKeys = append(table.PrimaryKeys, key.Name)
		}
	}
}

func GetColumns(db *schema.Database, table *schema.Table) {
	rows, err := db.Connection.Query("SELECT column_name, data_type from information_schema.columns WHERE table_catalog = $1 AND table_name = $2", db.Name, table.Name)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var colName string
		var colType string
		rows.Scan(&colName, &colType)
		c := &schema.Column{
			Name: colName,
			Type: colType,
		}
		table.Columns = append(table.Columns, c)
	}
}

func GetPrimaryKeys(db *schema.Database, tablename string) []*schema.Column {
	keys := []*schema.Column{}

	query := `SELECT a.attname, format_type(a.atttypid, a.atttypmod) AS data_type
						FROM   pg_index i
						JOIN   pg_attribute a ON a.attrelid = i.indrelid
																 AND a.attnum = ANY(i.indkey)
						WHERE  i.indrelid = $1::regclass
						AND    i.indisprimary`

	rows, err := db.Connection.Query(query, tablename)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var dataType string
		err := rows.Scan(&name, &dataType)
		if err != nil {
			fmt.Println(err)
		}
		c := &schema.Column{
			Name: name,
			Type: dataType,
		}
		keys = append(keys, c)
	}

	return keys
}
