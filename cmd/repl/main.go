package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jimmymuthoni/pesapal_rdbms_challenge/database"
	"github.com/jimmymuthoni/pesapal_rdbms_challenge/sql"
)

func main() {
	database.LoadSchemas()
	fmt.Println("RDBMS - Pesapal Challenge")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("db > ")
		input, _ := reader.ReadString('\n')
		tokens := sql.Parse(input)

		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {

		case "exit":
			return

		// ---------------- DATABASE COMMANDS ----------------
		case "CREATE":
			if len(tokens) >= 3 && tokens[1] == "DATABASE" {
				err := database.CreateDatabase(tokens[2])
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Database created:", tokens[2])
				}
				continue
			}

			// CREATE TABLE
			table := tokens[2]
			schema := database.TableSchema{
				Name: table,
				Columns: []database.Column{
					{Name: "id", Type: "INT"},
					{Name: "name", Type: "TEXT"},
				},
				PrimaryKey: "id",
				UniqueKeys: map[string]bool{},
			}
			database.CreateTable(schema)
			fmt.Println("Table created:", table)

		case "USE":
			err := database.UseDatabase(tokens[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Using database:", tokens[1])
			}

		case "SHOW":
			if tokens[1] == "DATABASES" {
				dbs, _ := database.ShowDatabases()
				for _, d := range dbs {
					fmt.Println(d)
				}
			}
			if tokens[1] == "TABLES" {
				tables, _ := database.ShowTables()
				for _, t := range tables {
					fmt.Println(t)
				}
			}

		// ---------------- CRUD COMMANDS ----------------

		case "INSERT":
			table := tokens[2]
			id, _ := strconv.Atoi(tokens[4])
			name := tokens[5]

			row := map[string]interface{}{
				"id":   id,
				"name": name,
			}

			err := database.Insert(table, row)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Inserted")
			}

		case "SELECT":
			table := tokens[3]
			rows, err := database.SelectAll(table)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, r := range rows {
				fmt.Println(r)
			}

		case "UPDATE":
			table := tokens[1]
			field := tokens[3]
			value := tokens[4]
			id, _ := strconv.Atoi(tokens[7])

			err := database.Update(table, id, field, value)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Updated")
			}

		case "DELETE":
			table := tokens[2]
			id, _ := strconv.Atoi(tokens[5])

			err := database.Delete(table, id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Deleted")
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
