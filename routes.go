package goctapus

import (
	"github.com/labstack/echo"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

var Routes map[string]Route

func GET(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	Server.GET(path, handler)
}

func POST(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	Server.POST(path, handler)
}

func PUT(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	Server.PUT(path, handler)
}

func DELETE(path string, handler echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	Server.DELETE(path, handler)
}

func AddStatic(path, file string) {
	Server.File(path, file)
}

func AddEndpoint(routeInfo Route) {
	switch routeInfo.Method {
	case "GET":
		GET(routeInfo.Path, routeInfo.Handler)
		break
	case "POST":
		POST(routeInfo.Path, routeInfo.Handler)
		break
	case "PUT":
		PUT(routeInfo.Path, routeInfo.Handler)
		break
	case "DELETE":
		DELETE(routeInfo.Path, routeInfo.Handler)
		break
	}
}
