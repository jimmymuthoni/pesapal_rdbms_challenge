package database

var Indexes = make(map[string]map[interface{}]int)

//function for indexing; indexing speeds up sql qiery execution
func BuildIndex(table string, pk string, rows[]map[string]interface{}){
	index := make(map[interface{}]int)
	for i, row := range rows {
		index[row[pk]] = i
	}
	Indexes[table] = index
}