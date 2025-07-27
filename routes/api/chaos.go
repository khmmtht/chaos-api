package api

import (
	"chaos-api/adapter"
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
)

func AddChaosRoutes(e *echo.Group) {
	chaosHandler := handler.NewChaosHandler(adapter.NewFileChaosAdapter())

	g := e.Group("/" + _const.ApiVersion + "/chaos")
	g.GET("/status/:service", chaosHandler.ChaosStatus)
	g.POST("/configure", chaosHandler.ChaosConfigure)
	g.POST("/trigger/:service", chaosHandler.ChaosTrigger)
	g.POST("/reset/:service", chaosHandler.ChaosReset)
}
