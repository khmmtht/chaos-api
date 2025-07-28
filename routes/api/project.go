package api

import (
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
)

func AddProjectRoutes(e *echo.Group) {
	handler := handler.NewProject()

	g := e.Group("/" + _const.ApiVersion + "/project")

	g.POST("", handler.NewProject)
	g.PATCH("", handler.UpdateProject)
	g.DELETE("", handler.RemoveProject)
	g.POST("/token", handler.GenApiKey)
}
