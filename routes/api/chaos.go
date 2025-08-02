package api

import (
	"chaos-api/adapter"
	_interface "chaos-api/adapter/interface"
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"os"
	"sync"
)

func AddChaosRoutes(e *echo.Group, store *sync.Map, client *mongo.Client) {
	var configAdapter _interface.ChaosConfigAdapter
	if os.Getenv("DRIVER") == "mongodb" {
		configAdapter = adapter.NewMongoDbChaosConfigAdapter(client)
	} else {
		configAdapter = adapter.NewMemoryChaosConfigAdapter(store)
	}

	chaosHandler := handler.NewChaosHandler(configAdapter)

	g := e.Group("/" + _const.ApiVersion + "/chaos")
	g.GET("/status/:service", chaosHandler.ChaosStatus)
	g.POST("/configure", chaosHandler.ChaosConfigure)
	g.POST("/trigger/:service", chaosHandler.ChaosTrigger)
	g.POST("/reset/:service", chaosHandler.ChaosReset)
}
