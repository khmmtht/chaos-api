package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"encoding/json"
	"os"
	"slices"
)

type FileChaosConfigAdapter struct {
	Path string
}

func NewFileChaosConfigAdapter() _interface.ChaosConfigAdapter {
	return FileChaosConfigAdapter{
		Path: "./data",
	}
}

func (a FileChaosConfigAdapter) UpsertChaosConfig(c *domain.ChaosConfig) error {
	configs, err := a.GetChaosConfigByProjectId(c.ProjectId)
	for i, config := range configs {
		if config.Name == c.Name {
			configs[i] = *c
		}
	}

	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath := a.Path + "/" + c.ProjectId + ".json"
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a FileChaosConfigAdapter) GetChaosConfigByProjectId(projectId string) ([]domain.ChaosConfig, error) {
	filePath := a.Path + "/" + projectId + ".json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var configs []domain.ChaosConfig
	err = json.Unmarshal(jsonData, &configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func (a FileChaosConfigAdapter) GetChaosConfigByService(projectId string, service string) (*domain.ChaosConfig, error) {
	configs, err := a.GetChaosConfigByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		if config.Name == service {
			return &config, nil
		}
	}

	return nil, nil
}

func (a FileChaosConfigAdapter) ResetConfig(projectId string, service string) error {
	configs, err := a.GetChaosConfigByProjectId(projectId)
	configs = slices.DeleteFunc(configs, func(c domain.ChaosConfig) bool {
		return c.Name == service
	})

	jsonData, err := json.Marshal(configs)
	if err != nil {
		return err
	}

	filePath := a.Path + "/" + projectId + ".json"
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
