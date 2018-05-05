package goctapus

import (
	"database/sql"
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

func ConnectDB(db_name string) {

	// Connects to a database and stores the connection to an object in the Databases Map
	Databases[db_name] = InitDB(Config.dbUser + ":" + Config.dbPass + "@tcp(" + Config.dbHost + ":" + Config.dbPort + ")/?charset=utf8")

	// Run the Database creation and USE queries
	sql := `CREATE DATABASE IF NOT EXISTS ` + db_name + `;
			USE ` + db_name + `;`
	executeSQLString(Databases[db_name], sql)

}

func executeSQLString(db *sql.DB, script string) {
	// split it into seperate queries
	queries := strings.Split(script, ";")

	// and execute them one by one
	// except for the last one which is expty because of the split
	for _, query := range queries[0 : len(queries)-1] {
		_, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}

func executeSQLFile(db *sql.DB, pathtofile string) {
	file, err := ioutil.ReadFile(pathtofile)

	if err != nil {
		panic(err)
	}

	executeSQLString(db, string(file))
}

// Migrate xecutes an SQL File against a specific DB
func Migrate(db *sql.DB, filename string) {
	executeSQLFile(db, filename)
}
