package goctapus

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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

func executeSQLFile(db *sql.DB, pathtofile string) {
	file, err := ioutil.ReadFile(pathtofile)

	if err != nil {
		panic(err)
	}

	// split it into seperate queries
	queries := strings.Split(string(file), ";")

	// and execute them one by one
	// except for the last one which is expty because of the split
	for _, query := range queries[0 : len(queries)-1] {
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}

// Migrate executes all the .sql files found inside the models directory
func Migrate(db *sql.DB) {

	executeSQLFile(db, "./models/main.sql")

	// Get all SQL files in models directory
	files := filesWithExtension("./models/", ".sql")

	// For each SQL file, read it and execute it
	for _, f := range files {
		if f.Name() == "main.sql" {
			continue
		}

		fmt.Println(f.Name())
		executeSQLFile(db, "./models/"+f.Name())
	}
}
