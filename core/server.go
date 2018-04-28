package goctapus

import (
	"database/sql"

	"github.com/Kamaropoulos/goctapus/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var appPort string
var dbHost string
var dbPort string
var dbUser string
var dbPass string

var db *sql.DB

var e *echo.Echo

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

func Init(args []string) {
	appPort, dbHost, dbPort, dbUser, dbPass = getArgs(args[1:])

	_ = appPort
	_ = dbHost
	_ = dbPort
	_ = dbUser
	_ = dbPass

	e = echo.New()
}

func Start() {
	e.Start(":" + appPort)
}

func AddEndpoints() {
	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))
}
