package goctapus

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_echo"
	"github.com/labstack/echo"

	Log "github.com/sirupsen/logrus"
)

type Route struct {
	Method   string
	Path     string
	Handler  echo.HandlerFunc
	Rate     float64
	_limiter *limiter.Limiter
}

var Routes map[string]Route

func GET(routeInfo Route) {
	//Check if Rate Limiter will be used for this route
	if routeInfo.Rate > 0 {
		// Register the route using the rate limiter middleware
		Server.GET(routeInfo.Path, routeInfo.Handler, tollbooth_echo.LimitHandler(routeInfo._limiter))
	} else {
		// Register the route without a rate limiter
		Server.GET(routeInfo.Path, routeInfo.Handler)
	}
}

func POST(routeInfo Route) {
	//Check if Rate Limiter will be used for this route
	if routeInfo.Rate > 0 {
		// Register the route using the rate limiter middleware
		Server.POST(routeInfo.Path, routeInfo.Handler, tollbooth_echo.LimitHandler(routeInfo._limiter))
	} else {
		// Register the route without a rate limiter
		Server.POST(routeInfo.Path, routeInfo.Handler)
	}
}

func PUT(routeInfo Route) {
	//Check if Rate Limiter will be used for this route
	if routeInfo.Rate > 0 {
		// Register the route using the rate limiter middleware
		Server.PUT(routeInfo.Path, routeInfo.Handler, tollbooth_echo.LimitHandler(routeInfo._limiter))
	} else {
		// Register the route without a rate limiter
		Server.PUT(routeInfo.Path, routeInfo.Handler)
	}
}

func DELETE(routeInfo Route) {
	//Check if Rate Limiter will be used for this route
	if routeInfo.Rate > 0 {
		// Register the route using the rate limiter middleware
		Server.DELETE(routeInfo.Path, routeInfo.Handler, tollbooth_echo.LimitHandler(routeInfo._limiter))
	} else {
		// Register the route without a rate limiter
		Server.DELETE(routeInfo.Path, routeInfo.Handler)
	}
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
	// Create Rate Limiter if Rate value is set
	if routeInfo.Rate > 0 {
		routeInfo._limiter = tollbooth.NewLimiter(routeInfo.Rate, nil)
	}

	// Store route into map and create it
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
