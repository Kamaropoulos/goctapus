package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/Kamaropoulos/goctapus/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func isValidPort(str string) (bool, string) {
	if i, err := strconv.Atoi(str); err == nil {
		if (1 <= i) && (i <= 65535) {
			return true, ""
		} else {
			return false, "A port number should be between 1 and 65535"
		}
	}
	return false, "Not a valid port number"
}

func isUsablePort(str string) (bool, string) {
	if res, err := isValidPort(str); !res {
		return res, err
	} else {
		ln, err1 := net.Listen("tcp", ":"+str)

		if err1 != nil {
			return false, err1.Error()
		}

		_ = ln.Close()

		return true, ""
	}
}

func getArgs(args []string) (string, string, string, string, string) {

	// Set default configuration values
	appPort := "8000"
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""

	switch argsCount := len(args); argsCount {
	case 1:
		// There was 1 argument passed, use it as
		// the port for the app to listen to
		if res, err := isUsablePort(args[0]); res {
			appPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

	case 2:
		// There were 2 arguments passed
		// Use them as DB username and password
		dbUser = args[0]
		dbPass = args[1]

	case 3:
		// There were 3 arguments passed
		// Use them as App Port, DB Username and Password
		if res, err := isUsablePort(args[0]); res {
			appPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

		dbUser = args[1]
		dbPass = args[2]

	case 4:
		// There were 4 arguments passed
		// Use them as App Port, DB Username, Password and Hostname/IP
		if res, err := isUsablePort(args[0]); res {
			appPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

		dbHost = args[3]
		dbUser = args[1]
		dbPass = args[2]

	case 5:
		// There were 5 arguments passed covering all configuration parameters
		// Use them as App Port, DB Username, Password, Hostname/IP and Port
		if res, err := isUsablePort(args[0]); res {
			appPort = args[0]
		} else {
			panic("\"" + args[0] + "\" can't be used as a port number: " + err)
		}

		dbHost = args[3]
		dbPort = args[4]
		dbUser = args[1]

		if res, err := isValidPort(args[2]); res {
			appPort = args[2]
		} else {
			panic("\"" + args[2] + "\" can't be used as a port number: " + err)
		}

		dbPass = args[2]
	}

	return appPort, dbHost, dbPort, dbUser, dbPass
}

func main() {
	appPort, dbHost, dbPort, dbUser, dbPass := getArgs(os.Args[1:])

	fmt.Println(appPort, dbHost, dbPort, dbUser, dbPass)

	//db := initDB(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/?charset=utf8")
	var db *sql.DB
	//migrate(db)

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
