package sql

import (
	"strings"
)

//function to process incoming sql queries
func Parse(query string) []string{
	query = strings.TrimSpace(query)
	query = strings.ReplaceAll(query,",","")
	return strings.Fields(query)
}