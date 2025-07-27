package api

import (
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
)

func AddChaosRoutes(e *echo.Group) {
	g := e.Group("/" + _const.ApiVersion + "/chaos")
	g.GET("/users/:id", handler.HelloWorld)
}
