package main

import (
	"database/sql"

	"./handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	db := initDB("<USERNAME>:<PASSWORD>@/?charset=utf8")

	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Start(":8000")
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

	sqlDB := `CREATE DATABASE IF NOT EXISTS todos`

	_, errDB := db.Exec(sqlDB)
	// Exit if something goes wrong with our SQL statement above
	if errDB != nil {
		panic(errDB)
	}

	_, err := db.Exec("USE todos")
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
