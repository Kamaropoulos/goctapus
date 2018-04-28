package goctapus

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the connection to the Database
func InitDB(dbString string) *sql.DB {
	db, err := sql.Open("mysql", dbString)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func Migrate(db *sql.DB) {

	sql := `CREATE DATABASE IF NOT EXISTS goapp;
			USE goapp;
			CREATE TABLE IF NOT EXISTS tasks(id INT NOT NULL AUTO_INCREMENT, name VARCHAR(50) NOT NULL, PRIMARY KEY (id));`

	queries := strings.Split(sql, ";")

	for _, query := range queries[0 : len(queries)-1] {
		// fmt.Println("--- ", query, " ---")
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}

}
