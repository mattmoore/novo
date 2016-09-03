package schema

import (
	"encoding/json"
	"io/ioutil"
)

func Parse(path string) *Schema {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	s := &Schema{}
	err = json.Unmarshal(fileContent, s)
	if err != nil {
		return nil
	}
	processSchema(s)
	return s
}

func processSchema(s *Schema) {
	for _, db := range s.Databases {
		for _, table := range db.Tables {
			table.Database = db
			for _, c := range table.Columns {
				c.Table = table
			}
		}
	}
}
