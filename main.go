package main

import (
	"database/sql"
	"os"

	goctapus "github.com/Kamaropoulos/goctapus/core"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args)

	//db := initDB(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/?charset=utf8")
	var db *sql.DB
	_ = db
	//migrate(db)

	goctapus.AddEndpoints()
	goctapus.Start()
}

func initDB(dbString string) *sql.DB {
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

func migrate(db *sql.DB) {

	sqlDB := `CREATE DATABASE IF NOT EXISTS goapp`

	_, errDB := db.Exec(sqlDB)
	// Exit if something goes wrong with our SQL statement above
	if errDB != nil {
		panic(errDB)
	}

	_, err := db.Exec("USE goapp")
	if err != nil {
		panic(err)
	}

	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INT NOT NULL AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
		PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
