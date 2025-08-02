package api

import (
	"chaos-api/adapter"
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddProjectRoutes(e *echo.Group, client *mongo.Client) {
	projectHandler := handler.NewProject(adapter.NewMongoDbProjectAdapter(client), adapter.NewMongoDbTokenAdapter(client))

	g := e.Group("/" + _const.ApiVersion + "/admin/project")

	g.POST("", projectHandler.NewProject)
	g.PATCH("", projectHandler.UpdateProject)
	g.DELETE("", projectHandler.RemoveProject)
	g.POST("/token", projectHandler.GenApiKey)
}
