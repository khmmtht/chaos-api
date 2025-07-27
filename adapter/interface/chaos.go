package _interface

import "chaos-api/domain"

type ChaosAdapter interface {
	UpsertChaosConfig(c *domain.ChaosConfig) error
	GetChaosConfig(token string, service string) (*domain.ChaosConfig, error)
	ResetConfig(token string, service string) error
}
