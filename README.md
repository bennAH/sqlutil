sqlutil
========
[godoc.org](https://godoc.org/github.com/bennAH/sqlutil)

golang sql util functions

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:

	go get github.com/bennAH/sqlutil

#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/bennAH/sqlutil"
```
### Raw Result Map Example
Saves needing to know columns to scan ahead of time.
```go

func main() {
	db, _ := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/db")
	rows, _ db.query("select * from table")
	rrm := rawResultMap(rows)
}
```
'rrm' -> a list of maps for each row result.

row result map -> string (column name) to []byte (column value)

### More to come
