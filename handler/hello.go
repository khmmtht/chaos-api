package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HelloWorld(c echo.Context) error {
	// User ID from path `users/:id`
	return c.String(http.StatusOK, "Hello, World!")
}
