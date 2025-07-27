package adapter

import (
	"chaos-api/domain"
	"encoding/json"
	"os"
)

type ChaosAdapter interface {
	CreateChaosConfig(c *domain.ChaosConfig) error
	GetChaosConfig(token string, service string) (*domain.ChaosConfig, error)
	ResetConfig(token string, service string) error
}

type FileChaosAdapter struct {
	Path string
}

func NewFileChaosAdapter() ChaosAdapter {
	return FileChaosAdapter{
		Path: "/data",
	}
}

func (a FileChaosAdapter) CreateChaosConfig(c *domain.ChaosConfig) error {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath := a.Path + "/" + c.Token + "_" + c.ServiceName + ".json"
	err = os.Truncate(filePath, 0)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (a FileChaosAdapter) GetChaosConfig(token string, service string) (*domain.ChaosConfig, error) {
	filePath := a.Path + "/" + token + "_" + service + ".json"
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
	err := os.Truncate(filePath, 0)
	if err != nil {
		return err
	}

	return nil
}
