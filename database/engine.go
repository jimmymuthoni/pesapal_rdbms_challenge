package database

import (
	"encoding/json"
	"errors"
	"os"
)

var CurrentDB = "default"
var Schemas = make(map[string]TableSchema)

// Load schemas and ensure default DB exists
func LoadSchemas() {
	os.MkdirAll("data/default", 0755)
	file, err := os.ReadFile("schemas.json")
	if err != nil {
		return
	}
	_ = json.Unmarshal(file, &Schemas)
}

func SaveSchemas() {
	bytes, _ := json.MarshalIndent(Schemas, "", "  ")
	_ = os.WriteFile("schemas.json", bytes, 0644)
}

func CreateDatabase(name string) error {
	return os.MkdirAll("data/"+name, 0755)
}

func ShowDatabases() ([]string, error) {
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
	return dbs, nil
}


func UseDatabase(name string) error {
	if _, err := os.Stat("data/" + name); err != nil {
		return err
	}
	CurrentDB = name
	return nil
}


func CreateTable(schema TableSchema) {
	Schemas[schema.Name] = schema
	SaveSchemas()
}

func ShowTables() ([]string, error) {
	path := "data/" + CurrentDB
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var tables []string
	for _, f := range files {
		if !f.IsDir() {
			tables = append(tables, f.Name())
		}
	}
	return tables, nil
}

func Update(table string, id int, field string, value interface{}) error {
	rows, err := ReadAllRows(table)
	if err != nil {
		return err
	}

	for _, row := range rows {
		if int(row["id"].(float64)) == id {
			row[field] = value
		}
	}
	return RewriteTable(table, rows)
}

func Delete(table string, id int) error {
	rows, err := ReadAllRows(table)
	if err != nil {
		return err
	}

	var updated []map[string]interface{}
	for _, row := range rows {
		if int(row["id"].(float64)) != id {
			updated = append(updated, row)
		}
	}
	return RewriteTable(table, updated)
}

func Join(left, right, leftKey, rightKey string) ([]map[string]interface{}, error) {
	lrows, _ := ReadAllRows(left)
	rrows, _ := ReadAllRows(right)

	var result []map[string]interface{}
	for _, l := range lrows {
		for _, r := range rrows {
			if l[leftKey] == r[rightKey] {
				row := make(map[string]interface{})
				for k, v := range l {
					row[left+"."+k] = v
				}
				for k, v := range r {
					row[right+"."+k] = v
				}
				result = append(result, row)
			}
		}
	}
	return result, nil
}

func Insert(table string, row map[string]interface{}) error {
	schema := Schemas[table]

	rows, _ := ReadAllRows(table)
	for _, r := range rows {
		if r[schema.PrimaryKey] == row[schema.PrimaryKey] {
			return errors.New("primary key violation")
		}
	}
	return InsertRow(table, row)
}


func SelectAll(table string) ([]map[string]interface{}, error) {
	return ReadAllRows(table)
}
