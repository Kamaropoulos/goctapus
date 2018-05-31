package goctapus

import (
	"database/sql"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	Log "github.com/sirupsen/logrus"
)

// InitDB initializes the connection to the Database
func InitDB(dbString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbString)

	// Here we check for any db errors then exit
	if err != nil {
		return nil, err
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		return nil, nil
	}

	// Ping to check if the server is up
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err;
	}

	return db, nil
}

func ConnectDB(db_name string) {

	Log.Debug("Connecting to Database server on " + Config.dbHost + ":" + Config.dbPort + "...")

	// Connects to a database and stores the connection to an object in the Databases Map
	db, err := InitDB(Config.dbUser + ":" + Config.dbPass + "@tcp(" + Config.dbHost + ":" + Config.dbPort + ")/?charset=utf8")

	if (err != nil) || (db == nil) {
		Log.Fatal("Could not connect to the Database server: " + err.Error())
	}

	Databases[db_name] = db

	Log.Debug("Connected to Database server.")

	Log.Debug("Creating DB \"" + db_name + "\" if it doesn't already exist...")

	// Run the Database creation and USE queries
	sql := `CREATE DATABASE IF NOT EXISTS ` + db_name + `;
			USE ` + db_name + `;`
	err = executeSQLString(Databases[db_name], sql)
	if err != nil {
		Log.Fatal("An error occured while creating DB \"" + db_name + "\"")
	}

}

func executeSQLString(db *sql.DB, script string) error {
	// split it into seperate queries
	queries := strings.Split(script, ";")

	// and execute them one by one
	// except for the last one which is expty because of the split
	for _, query := range queries[0 : len(queries)-1] {
		_, err := db.Exec(query)
		if err != nil {
			return err;
		}
	}
	return nil;
}

func executeSQLFile(db *sql.DB, pathtofile string) error {
	file, err := ioutil.ReadFile(pathtofile)

	if err != nil {
		return err
	}

	return executeSQLString(db, string(file))
}

// Migrate xecutes an SQL File against a specific DB
func Migrate(db *sql.DB, filename string) {

	Log.Debug("Exexuting SQL file \"" + filename + "\"...")

	err := executeSQLFile(db, filename)

	if err != nil {
		Log.Fatal("Executing " + filename + " failed: " + err.Error())
	}

	Log.Debug("SQL file \"" + filename + "\" was executed succesfully.")
}
