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
	adapter _interface.ChaosAdapter
}

func NewChaosHandler(adapter _interface.ChaosAdapter) *ChaosHandler {
	return &ChaosHandler{
		adapter: adapter,
	}
}

type ChaosConfigRequest struct {
	ServiceName string      `json:"service_name"`
	Mode        domain.Mode `json:"mode"`
	Value       string      `json:"value"`
	Response    string      `json:"response"`
}

func (h *ChaosHandler) ChaosStatus(c echo.Context) error {
	t := c.Request().Header.Get("project-id")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api ProjectId required")
	}

	config, err := h.adapter.GetChaosConfigByService(t, c.Param("service"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, config)
}

func (h *ChaosHandler) ChaosConfigure(c echo.Context) error {
	config := new(domain.ChaosConfig)
	if err := c.Bind(config); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.adapter.UpsertChaosConfig(config)
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
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	switch config.Mode {
	case domain.Latency:
		delay, err := strconv.Atoi(config.Value)
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)

		return c.JSON(http.StatusOK, config.Response)
	case domain.Response:
		code, err := strconv.Atoi(config.Value)
		if err != nil {
			return err
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
