package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jimmymuthoni/pesapal_rdbms_challenge/database"
)


type User struct {
	ID 	  int `json:"id"`
	Name  string `jaon:"name"`
}

func main(){
	//loading the schema
	database.LoadSchemas()
	
	if _, ok := database.Schemas["users"]; !ok {
		database.CreateTable(database.TableSchema{
			Name: "users",
			Columns: []database.Column{
				{Name: "id", Type: "INT"},
				{Name: "name", Type: "TEXT"},

			},
			PrimaryKey: "id",
			UniqueKeys: map[string]bool{},
		})
	}

	http.HandleFunc("/users", userHandler)
	http.HandleFunc("/users/",userByIDHandler)

	log.Println("Trival web app to interact with custom RDBMS..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// POST /users -> Create user
// GET /users  -> Read all users
func userHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{

	case http.MethodPost:
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w,"Invalid JSON", http.StatusBadRequest)
			return
		}
		row := map[string]interface{}{
			"id": user.ID,
			"name": user.Name,
		}

		if err := database.Insert("users", row); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created"))

	case http.MethodGet:
		rows, err := database.SelectAll("users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(rows)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


// Delete /users/{id} -> Delete user
func userByIDHandler(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	
	case http.MethodDelete:
		rows, err := database.SelectAll("users")
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var updated []map[string]interface{}
		for _, row := range rows {
			if int (row["id"].(float64)) != id {
				updated = append(updated, row)
			}
		}

		database.RewriteTable("users", updated)
		w.Write([]byte("User deleted"))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}