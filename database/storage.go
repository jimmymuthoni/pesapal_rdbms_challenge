package database

import (
	"bufio"
	"encoding/json"
	"os"
	
)

//function to insert row to table in database
func InsertRow(table string, row map[string]interface{}) error {
	file, err := os.OpenFile("data/" + CurrentDB + "/" + table + ".tbl", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	bytes, _ := json.Marshal(row)
	writer.Write(bytes)
	writer.WriteString("\n")
	return writer.Flush()

}

//function to read all rows in a table
func ReadAllRows(table string) ([]map[string]interface{}, error){
	file, err := os.Open("data/"+table+".tbl")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var rows []map[string]interface{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var row map[string]interface{}
		json.Unmarshal(scanner.Bytes(), &row)
		rows = append(rows, row)

	}
	return  rows, nil
}


func RewriteTable(table string, rows []map[string]interface{}) error {
	file, err := os.Create("data/" + table+ ".tbl")
	if err != nil {
		return  err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, row := range rows {
		bytes, _ := json.Marshal(row)
		writer.Write(bytes)
		writer.WriteString("\n")
	}
	return writer.Flush()
}