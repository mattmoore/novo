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
	s.Compile()
	return s
}
