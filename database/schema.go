package database

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TableSchema struct{
	Name		string   		`json:"name"`
	Columns 	[]Column 		`json:"columns"`
	PrimaryKey 	string 			`json:"primary_key"`
	UniqueKeys 	map[string]bool `json:"unique_keys"`
}