package database

import (
	"encoding/json"
	"errors"
	"os"
)

var CurrentDM = "default"
var Schemas = make(map[string]TableSchema)

// loading schema
func LoadSchemas(){
	os.MkdirAll("data/default", 0755)
	file, err := os.ReadFile("schemas.json")
	if err != nil {
		return 
	}
	json.Unmarshal(file, &Schemas)
}

//save the schema
func SaveSchemas(){
	bytes, _ := json.MarshalIndent(Schemas,""," ")
	os.WriteFile("schemas.json",bytes,0644)
}

func CreateTable(schema TableSchema){
	Schemas[schema.Name] = schema
	SaveSchemas()
}

func Insert(table string, row map[string]interface{}) error {
	schema := Schemas[table]

	rows, _ := ReadAllRows(table)
	for _, r := range rows {
		if r[schema.PrimaryKey] == row[schema.PrimaryKey]{
			return errors.New("Primary key violation")
		}
	}
	return InsertRow(table, row)
}

func SelectAll(table string) ([]map[string]interface{}, error){
	return ReadAllRows(table)
}

