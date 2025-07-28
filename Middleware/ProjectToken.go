package Middleware

import (
	_interface "chaos-api/adapter/interface"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProjectTokenMiddleware struct {
	TokenAdapter _interface.TokenAdapter
}

func NewProjectTokenMiddleware(tokenAdapter _interface.TokenAdapter) *ProjectTokenMiddleware {
	return &ProjectTokenMiddleware{TokenAdapter: tokenAdapter}
}

func (b *ProjectTokenMiddleware) Handler() echo.MiddlewareFunc {
	// TODO: Implement Cache or Add logic validate at handler
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Api-Token")
			projectId := c.Request().Header.Get("Project-Id")
			if token == "" || projectId == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Api-Token or Project-Id header")
			}

			count, err := b.TokenAdapter.Count(projectId, token)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			if count <= 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
			}

			return next(c)
		}
	}
}
