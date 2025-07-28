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
	result, err := p.ProjectAdapter.CreateProject(c.FormValue("name"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func (p *ProjectHandler) UpdateProject(c echo.Context) error {
	err := p.ProjectAdapter.UpdateProject(c.FormValue("project-id"), c.FormValue("name"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Project Updated")
}

func (p *ProjectHandler) RemoveProject(c echo.Context) error {
	// TODO: Remove cascade
	err := p.ProjectAdapter.DeleteProject(c.FormValue("project-id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Project Deleted")
}

func (p *ProjectHandler) GenApiKey(c echo.Context) error {
	result, err := p.TokenAdapter.GenerateToken(c.FormValue("project-id"), c.FormValue("name"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}
