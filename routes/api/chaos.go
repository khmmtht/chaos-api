package api

import (
	"chaos-api/adapter"
	_const "chaos-api/const"
	"chaos-api/handler"
	"chaos-api/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AddChaosRoutes(e *echo.Group, client *mongo.Client) {
	chaosHandler := handler.NewChaosHandler(adapter.NewMongoDbChaosConfigAdapter(client))
	middleware := middleware.NewProjectTokenMiddleware(adapter.NewMongoDbTokenAdapter(client))

	g := e.Group("/"+_const.ApiVersion+"/chaos", middleware.Handler())
	g.GET("/status/:service", chaosHandler.ChaosStatus)
	g.POST("/configure", chaosHandler.ChaosConfigure)
	g.POST("/trigger/:service", chaosHandler.ChaosTrigger)
	g.POST("/reset/:service", chaosHandler.ChaosReset)
}
