package _interface

import "chaos-api/domain"

type ProjectAdapter interface {
	GetProjects() ([]domain.Project, error)
	CreateProject(name string) (*domain.Project, error)
	UpdateProject(projectId, name string) error
	DeleteProject(projectId string) error
}
