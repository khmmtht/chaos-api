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
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath := a.Path + "/" + c.ProjectId + "_" + c.Name + ".json"
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a FileChaosAdapter) GetChaosConfigByProjectId(projectId string) (*domain.ChaosConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (a FileChaosAdapter) GetChaosConfigByService(projectId string, service string) (*domain.ChaosConfig, error) {
	filePath := a.Path + "/" + projectId + "_" + service + ".json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config domain.ChaosConfig
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (a FileChaosAdapter) ResetConfig(token string, service string) error {
	filePath := a.Path + "/" + token + "_" + service + ".json"
	err := os.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		return err
	}

	return nil
}
