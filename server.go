package goctapus

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	Log "github.com/sirupsen/logrus"
)

var Config Configuration

type Configuration struct {
	appPort string
	dbHost  string
	dbPort  string
	dbUser  string
	dbPass  string
}

var Databases map[string]*sql.DB

var Server *echo.Echo

func getArgs(args []string) Configuration {

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

	conf := Configuration{appPort, dbHost, dbPort, dbUser, dbPass}

	return conf
}

func Init(args []string, logLevel string) {

	InitLogger(logLevel)

	Log.Info("Initializing Goctapus...")

	Config = getArgs(args[1:])

	Log.WithFields(Log.Fields{
		"appPort": Config.appPort,
		"dbHost":  Config.dbHost,
		"dbPort":  Config.dbPort,
		"dbUser":  Config.dbUser,
		"dbPass":  Config.dbPass,
	}).Debug("Current server configuration:")

	Databases = make(map[string]*sql.DB)

	Server = echo.New()

	Log.Debug("Goctapus Initialization done.")
}

func Start() {
	Log.Info("Starting up web server...")
	Server.Start(":" + Config.appPort)
	Log.Debug("Web Server started succesfully.")
}
