package api

import (
	"chaos-api/adapter"
	_interface "chaos-api/adapter/interface"
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"os"
)

func AddProjectRoutes(e *echo.Group, client *mongo.Client) {
	var projectAdapter _interface.ProjectAdapter
	var tokenAdapter _interface.TokenAdapter
	if os.Getenv("DRIVER") == "mongodb" {
		projectAdapter = adapter.NewMongoDbProjectAdapter(client)
		tokenAdapter = adapter.NewMongoDbTokenAdapter(client)
	} else {
		projectAdapter = adapter.NewMemoryProjectAdapter()
		tokenAdapter = adapter.NewMemoryTokenAdapter()
	}

	projectHandler := handler.NewProject(projectAdapter, tokenAdapter)

	g := e.Group("/" + _const.ApiVersion + "/admin/project")

	g.POST("", projectHandler.NewProject)
	g.PATCH("", projectHandler.UpdateProject)
	g.DELETE("", projectHandler.RemoveProject)
	g.POST("/token", projectHandler.GenApiKey)
}
