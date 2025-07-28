package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"encoding/json"
	"os"
)

type FileChaosAdapter struct {
	Path string
}

func NewFileChaosAdapter() _interface.ChaosAdapter {
	return FileChaosAdapter{
		Path: "./data",
	}
}

func (a FileChaosAdapter) UpsertChaosConfig(c *domain.ChaosConfig) error {
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

func (a FileChaosAdapter) GetChaosConfigByProjectId(projectId string) ([]domain.ChaosConfig, error) {
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

func (a FileChaosAdapter) GetChaosConfigByService(projectId string, service string) (*domain.ChaosConfig, error) {
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

func (a FileChaosAdapter) ResetConfig(token string, service string) error {
	filePath := a.Path + "/" + token + ".json"
	err := os.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		return err
	}

	return nil
}
