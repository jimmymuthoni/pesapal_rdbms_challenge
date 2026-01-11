package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jimmymuthoni/pesapal_rdbms_challenge/database"
	"github.com/jimmymuthoni/pesapal_rdbms_challenge/sql"

)

func main(){
	database.LoadSchemas()
	fmt.Println("RDBMS -Pesapal Challenge")
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

		case "CREATE":
			table := tokens[2]
			schema := database.TableSchema{
				Name: table,
				Columns: []database.Column{
					{Name: "id", Type: "INT"},
					{Name: "name", Type: "INT"},
				},
				PrimaryKey: "id",
				UniqueKeys: map[string]bool{},
			}
			database.CreateTable(schema)
			fmt.Println("Table created:", table)

		case "INSERT":
			table := tokens[2]
			id, _ := strconv.Atoi(tokens[4])
			name := tokens[5]

			row := map[string]interface{}{
				"id": id,
				"name":name,
			}

			err := database.Insert(table, row)
			if err != nil{
				fmt.Println(err)		
			} else{
				fmt.Println("Inserted")
			}
		
		case "SELECT":
			table := tokens[3]
			rows,_ := database.SelectAll(table)
			for _, r := range rows {
				fmt.Println(r)
			}
		default:
			fmt.Println("Unknown command")
		}	
	
	}
}