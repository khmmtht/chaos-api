package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type MemoryProjectAdapter struct {
	Store  *sync.Map
	prefix string
}

func NewMemoryProjectAdapter(store *sync.Map) _interface.ProjectAdapter {
	return &MemoryProjectAdapter{
		Store:  store,
		prefix: "project_",
	}
}

func (a *MemoryProjectAdapter) generateKey(key string) string {
	return a.prefix + key
}

func (a *MemoryProjectAdapter) GetProjects() ([]domain.Project, error) {
	projects := make([]domain.Project, 0)
	a.Store.Range(func(key, value interface{}) bool {
		projects = append(projects, value.(domain.Project))
		return true
	})

	return projects, nil
}

func (a *MemoryProjectAdapter) CreateProject(name string) (*domain.Project, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	project := domain.Project{
		Id:   id.String(),
		Name: name,
	}
	a.Store.Store(a.generateKey(id.String()), project)

	return &project, nil
}

func (a *MemoryProjectAdapter) UpdateProject(projectId, name string) error {
	val, ok := a.Store.Load(a.generateKey(projectId))
	if !ok {
		return errors.New("project not found")
	}

	project := val.(domain.Project)
	project.Name = name
	a.Store.Store(a.generateKey(projectId), project)

	return nil
}

func (a *MemoryProjectAdapter) DeleteProject(projectId string) error {
	a.Store.Delete(a.generateKey(projectId))
	return nil
}
