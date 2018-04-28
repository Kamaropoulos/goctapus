package goctapus

import (
	"database/sql"
	"io/ioutil"
	"log"
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

// Migrate executes all the .sql files found inside the models directory
func Migrate(db *sql.DB) {

	// Get all files in models directory
	files, err := ioutil.ReadDir("./models/")
	if err != nil {
		log.Fatal(err)
	}

	// For each SQL file, read it and execute it
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".sql") {
			// f is an SQL file, read it
			file, err := ioutil.ReadFile("./models/" + f.Name())

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
	}
}
