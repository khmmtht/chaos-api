package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"errors"
	"sync"
)

type MemoryChaosConfigAdapter struct {
	Store *sync.Map
}

func NewMemoryChaosConfigAdapter(store *sync.Map) _interface.ChaosConfigAdapter {
	return &MemoryChaosConfigAdapter{
		Store: store,
	}
}

func generateKey(projectId, name string) string {
	return projectId + "_" + name
}

func (a *MemoryChaosConfigAdapter) UpsertChaosConfig(c *domain.ChaosConfig) error {
	a.Store.Store(generateKey(c.ProjectId, c.Name), c)

	return nil
}

func (a *MemoryChaosConfigAdapter) GetChaosConfigByProjectId(projectId string) ([]domain.ChaosConfig, error) {
	configs := make([]domain.ChaosConfig, 0)
	a.Store.Range(func(key, value any) bool {
		strKey, ok := key.(string)
		if !ok {
			return true // skip non-string keys
		}
		if len(strKey) >= len(projectId) && strKey[:len(projectId)] == projectId {
			configs = append(configs, value.(domain.ChaosConfig))
		}
		return true
	})

	return configs, nil
}

func (a *MemoryChaosConfigAdapter) GetChaosConfigByService(projectId string, service string) (*domain.ChaosConfig, error) {
	result, ok := a.Store.Load(generateKey(projectId, service))
	if !ok {
		return nil, errors.New("key does not exist")
	}

	config := result.(domain.ChaosConfig)
	return &config, nil
}

func (a *MemoryChaosConfigAdapter) ResetConfig(projectId string, service string) error {
	a.Store.Delete(generateKey(projectId, service))
	return nil
}
