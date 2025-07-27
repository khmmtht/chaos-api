package _interface

import "chaos-api/domain"

type ChaosAdapter interface {
	UpsertChaosConfig(c *domain.ChaosConfig) error
	GetChaosConfig(projectId string, service string) (*domain.ChaosConfig, error)
	ResetConfig(projectId string, service string) error
}
