package api

import (
	_const "chaos-api/const"
	"chaos-api/handler"
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddSimulateRoutes(e *echo.Group) {
	g := e.Group("/" + _const.ApiVersion + "/simulate")

	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions, http.MethodTrace, http.MethodConnect, http.MethodHead}

	for _, met := range methods {
		g.Add(met, "/delay/:ms", handler.SimulateDelay)
		g.Add(met, "/error/:code", handler.SimulateError)
	}
}
