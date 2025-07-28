package api

import (
	"chaos-api/adapter"
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddProjectRoutes(e *echo.Group, client *mongo.Client) {
	handler := handler.NewProject(adapter.NewMongoDbProjectAdapter(client), adapter.NewMongoDbTokenAdapter(client))

	g := e.Group("/" + _const.ApiVersion + "/project")

	g.POST("", handler.NewProject)
	g.PATCH("", handler.UpdateProject)
	g.DELETE("", handler.RemoveProject)
	g.POST("/token", handler.GenApiKey)
}
