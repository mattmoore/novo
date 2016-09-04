package schema

import (
	"encoding/json"
	"io/ioutil"
)

func Parse(path string) *Database {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	db := &Database{}
	err = json.Unmarshal(fileContent, db)
	if err != nil {
		return nil
	}
	db.Compile()
	return db
}
