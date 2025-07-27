package handler

import (
	"chaos-api/adapter"
	"chaos-api/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ChaosHandler struct {
	adapter adapter.ChaosAdapter
}

func NewChaosHandler(adapter adapter.ChaosAdapter) *ChaosHandler {
	return &ChaosHandler{
		adapter: adapter,
	}
}

type ChaosTriggerRequest struct {
	UserId      string `json:"user_id"`
	ServiceName string `json:"service_name"`
}

func (h *ChaosHandler) ChaosStatus(c echo.Context) error {
	t := c.Request().Header.Get("api-token")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api Token required")
	}

	config, err := h.adapter.GetChaosConfig(t, c.Param("service"))
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

	err := h.adapter.CreateChaosConfig(config)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.String(http.StatusCreated, "Created Chaos Config")
}

func (h *ChaosHandler) ChaosTrigger(c echo.Context) error {
	t := c.Request().Header.Get("api-token")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api Token required")
	}

	// Main logic
	config, err := h.adapter.GetChaosConfig(t, c.Param("service"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	switch config.Mode {
	case domain.Latency:
		return c.JSON(http.StatusOK, config)
	case domain.Error:
		code, err := strconv.Atoi(config.Value)
		if err != nil {
			return err
		}

		return c.String(code, config.Response)
	default:
		return c.JSON(http.StatusBadRequest, "Invalid Chaos Mode")
	}
}

func (h *ChaosHandler) ChaosReset(c echo.Context) error {
	t := c.Request().Header.Get("api-token")
	if len(t) == 0 {
		return c.JSON(http.StatusBadRequest, "Api Token required")
	}

	err := h.adapter.ResetConfig(t, c.Param("service"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, "success")
}
