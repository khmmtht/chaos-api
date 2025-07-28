package handler

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type ChaosHandler struct {
	adapter _interface.ChaosConfigAdapter
}

func NewChaosHandler(adapter _interface.ChaosConfigAdapter) *ChaosHandler {
	return &ChaosHandler{
		adapter: adapter,
	}
}

type ChaosConfigRequest struct {
	Name     string      `json:"name"`
	Mode     domain.Mode `json:"mode"`
	Value    string      `json:"value"`
	Response string      `json:"response"`
}

func (h *ChaosHandler) ChaosStatus(c echo.Context) error {
	t := c.Request().Header.Get("project-id")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api ProjectId required")
	}

	config, err := h.adapter.GetChaosConfigByService(t, c.Param("service"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, config)
}

func (h *ChaosHandler) ChaosConfigure(c echo.Context) error {
	t := c.Request().Header.Get("project-id")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api ProjectId required")
	}

	config := new(ChaosConfigRequest)
	if err := c.Bind(config); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.adapter.UpsertChaosConfig(&domain.ChaosConfig{
		ProjectId: t,
		Name:      config.Name,
		Mode:      config.Mode,
		Value:     config.Value,
		Response:  config.Response,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, "Created Chaos Config")
}

func (h *ChaosHandler) ChaosTrigger(c echo.Context) error {
	t := c.Request().Header.Get("project-id")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api ProjectId required")
	}

	// Main logic
	config, err := h.adapter.GetChaosConfigByService(t, c.Param("service"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	switch config.Mode {
	case domain.Latency:
		delay, err := strconv.Atoi(config.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)

		return c.JSON(http.StatusOK, config.Response)
	case domain.Response:
		code, err := strconv.Atoi(config.Value)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.String(code, config.Response)
	default:
		return c.JSON(http.StatusInternalServerError, "Invalid Chaos Mode")
	}
}

func (h *ChaosHandler) ChaosReset(c echo.Context) error {
	t := c.Request().Header.Get("project-id")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api ProjectId required")
	}

	err := h.adapter.ResetConfig(t, c.Param("service"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, "success")
}
