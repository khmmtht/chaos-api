package api

import (
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
)

func AddGlobalRoutes(e *echo.Group) {
	g := e.Group("/" + _const.ApiVersion + "/global")
	g.GET("/config", handler.HelloWorld)
}
