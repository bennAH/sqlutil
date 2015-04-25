package sqlutil

import (
	"database/sql"
)

// Creates a result holder for a query result as an array of raw bytes
func ResultHolder(rows *sql.Rows) []interface{} {
	cols, _ := rows.Columns()
	l := len(cols)
	result := make([]interface{}, l, l)
	for i, _ := range cols {
		result[i] = new(sql.RawBytes)
	}
	return result
}

type rawResult map[string][]byte

// Converts a query results row for each row resultt as a map of column to raw bytes
func RawResultMap(rows *sql.Rows) []rawResult {
	cols, _ := rows.Columns()
	result := ResultHolder(rows)
	results := make([]rawResult, 0)
	c := 0
	for rows.Next() {
		c++
		resultMap := make(map[string][]byte)
		rows.Scan(result...)
		for i, v := range result {
			f := v.(*sql.RawBytes)
			resultMap[cols[i]] = (*f)[:]
		}
		results = append(results, resultMap)
	}

	return results
}
