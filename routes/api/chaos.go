package api

import (
	"chaos-api/adapter"
	_interface "chaos-api/adapter/interface"
	_const "chaos-api/const"
	"chaos-api/handler"
	"chaos-api/middleware"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"os"
	"sync"
)

func AddChaosRoutes(e *echo.Group, store *sync.Map, client *mongo.Client) {
	var configAdapter _interface.ChaosConfigAdapter
	var tokenAdapter _interface.TokenAdapter
	if os.Getenv("DRIVER") == "mongodb" {
		configAdapter = adapter.NewMongoDbChaosConfigAdapter(client)
		tokenAdapter = adapter.NewMongoDbTokenAdapter(client)
	} else {
		configAdapter = adapter.NewMemoryChaosConfigAdapter(store)
		tokenAdapter = adapter.NewMemoryTokenAdapter(store)
	}

	chaosHandler := handler.NewChaosHandler(configAdapter)
	tokenMiddleware := middleware.NewProjectTokenMiddleware(tokenAdapter)

	g := e.Group("/"+_const.ApiVersion+"/chaos", tokenMiddleware.Handler())
	g.GET("/status/:service", chaosHandler.ChaosStatus)
	g.POST("/configure", chaosHandler.ChaosConfigure)
	g.POST("/trigger/:service", chaosHandler.ChaosTrigger)
	g.POST("/reset/:service", chaosHandler.ChaosReset)
}
