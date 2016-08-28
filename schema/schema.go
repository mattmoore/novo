package schema

type Schema struct {
	Databases map[string]Database
}

type Database struct {
	Tables map[string]Table
}

type Table struct {
	Columns map[string]Column
}
