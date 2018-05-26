package goctapus

import (
	"github.com/labstack/echo"

	Log "github.com/sirupsen/logrus"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

var Routes map[string]Route

func GET(routeInfo Route) {
	Server.GET(routeInfo.Path, routeInfo.Handler)
}

func POST(routeInfo Route) {
	Server.POST(routeInfo.Path, routeInfo.Handler)
}

func PUT(routeInfo Route) {
	Server.PUT(routeInfo.Path, routeInfo.Handler)
}

func DELETE(routeInfo Route) {
	Server.DELETE(routeInfo.Path, routeInfo.Handler)
}

func AddStatic(path, file string) {
	Server.File(path, file)
}

func AddEndpoint(routeInfo Route) {
	//Generate route descriptor string
	routeName := routeInfo.Method + ":" + routeInfo.Path
	//Check if the route already exists
	if _, ok := Routes[routeName]; ok {
		// Route already exists
		// Just return with a message
		Log.Warning("Could not add endpoint " + routeName + ". Route already exists!")
		return
	}

	// Route doesn't exist
	// Store it to the map and create it
	Routes[routeName] = routeInfo

	switch routeInfo.Method {
	case "GET":
		GET(routeInfo)
		break
	case "POST":
		POST(routeInfo)
		break
	case "PUT":
		PUT(routeInfo)
		break
	case "DELETE":
		DELETE(routeInfo)
		break
	}
}
