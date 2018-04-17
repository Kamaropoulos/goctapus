package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Kamaropoulos/go-echo-vue-mysql/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func getArgs(args []string) (string, string, string, string, string) {
	var appPort string
	var dbHost string
	var dbPort string
	var dbUser string
	var dbPass string

	switch argsCount := len(args); argsCount {
	case 0:
		// No arguments were passed
		// Using default configuration values
		appPort = "8000"
		dbHost = "localhost"
		dbPort = "3306"
		dbUser = "root"
		dbPass = ""
	case 1:
		// There was 1 argument passed, use it as
		// the port for the app to listen to
		appPort = args[0]
		dbHost = "localhost"
		dbPort = "3306"
		dbUser = "root"
		dbPass = ""
	case 2:
		// There were 2 arguments passed
		// Use them as DB username and password
		appPort = "8000"
		dbHost = "localhost"
		dbPort = "3306"
		dbUser = args[0]
		dbPass = args[1]
	case 3:
		// There were 3 arguments passed
		// Use them as App Port, DB Username and Password
		appPort = args[0]
		dbHost = "localhost"
		dbPort = "3306"
		dbUser = args[1]
		dbPass = args[2]
	case 4:
		// There were 4 arguments passed
		// Use them as App Port, DB Username, Password and Hostname/IP
		appPort = args[0]
		dbHost = args[3]
		dbPort = "3306"
		dbUser = args[1]
		dbPass = args[2]
	case 5:
		// There were 5 arguments passed covering all configuration parameters
		// Use them as App Port, DB Username, Password, Hostname/IP and Port
		appPort = args[0]
		dbHost = args[3]
		dbPort = args[4]
		dbUser = args[1]
		dbPass = args[2]
	}

	return appPort, dbHost, dbPort, dbUser, dbPass
}

func main() {
	appPort, dbHost, dbPort, dbUser, dbPass := getArgs(os.Args[1:])

	fmt.Println(appPort, dbHost, dbPort, dbUser, dbPass)

	db := initDB(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/?charset=utf8")

	migrate(db)

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	e.Start(":" + appPort)
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
