package database

import (
	"encoding/json"
	"errors"
	"os"
)

var CurrentDB = "default"
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

//function to create db
func CreateDatabase(name string) error{
	return os.MkdirAll("data/"+name, 0755)
}

//showing databases
func ShowDatabases() ([]string, error){
	files, err := os.ReadDir("data")
	if err != nil {
		return nil, err
	}

	var dbs []string
	for _, f := range files {
		if f.IsDir() {
			dbs = append(dbs, f.Name())
		}
	}
	return dbs,nil
}

//selcting database to use 
func UseDatabase(name string) error {
	if _, err := os.Stat("/data/" + name); err != nil {
		return err
	}
	CurrentDB = name
	return  nil
}

//creating table in slected database
func CreateTable(schema TableSchema){
	Schemas[schema.Name] = schema
	SaveSchemas()
}


//showing the tables avaiable in database
func ShowTables() ([]string, error){
	path := "data/"+CurrentDB
	files, err := os.ReadDir(path)
	if err != nil {
		return  nil, err
	}

	var tables []string
	for _, f := range files{
		if !f.IsDir(){
			tables = append(tables, f.Name())
		}
	}
	return tables, nil
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

