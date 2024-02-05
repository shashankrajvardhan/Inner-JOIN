package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	connection()
	join()
	// r := mux.NewRouter()
	// rout(r)
	// fmt.Println("running in port:8000")
	// log.Fatal(http.ListenAndServe(":8000", r))

}
