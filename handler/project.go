package handler

import (
	_interface "chaos-api/adapter/interface"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProjectHandler struct {
	ProjectAdapter _interface.ProjectAdapter
	TokenAdapter   _interface.TokenAdapter
}

func NewProject(projectAdapter _interface.ProjectAdapter, tokenAdapter _interface.TokenAdapter) *ProjectHandler {
	return &ProjectHandler{
		ProjectAdapter: projectAdapter,
		TokenAdapter:   tokenAdapter,
	}
}

func (p *ProjectHandler) NewProject(c echo.Context) error {
	var data map[string]interface{}
	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := p.ProjectAdapter.CreateProject(data["name"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (p *ProjectHandler) UpdateProject(c echo.Context) error {
	var data map[string]interface{}
	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = p.ProjectAdapter.UpdateProject(data["project_id"].(string), data["name"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Project Updated")
}

func (p *ProjectHandler) RemoveProject(c echo.Context) error {
	// TODO: Remove cascade
	var data map[string]interface{}
	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = p.ProjectAdapter.DeleteProject(data["project_id"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Project Deleted")
}

func (p *ProjectHandler) GenApiKey(c echo.Context) error {
	var data map[string]interface{}
	err := c.Bind(&data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	result, err := p.TokenAdapter.GenerateToken(data["project_id"].(string), data["name"].(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
