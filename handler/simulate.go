package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

func SimulateDelay(c echo.Context) error {
	d, err := strconv.Atoi(c.Param("ms"))
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(d) * time.Millisecond)
	return c.String(http.StatusOK, "OK")
}

func SimulateError(c echo.Context) error {
	co := c.Param("code")
	code, err := strconv.Atoi(co)
	if err != nil {
		return err
	}

	return c.String(code, co)
}
