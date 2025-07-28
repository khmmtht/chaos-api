package _interface

import "chaos-api/domain"

type TokenAdapter interface {
	GetTokens() ([]domain.Token, error)
	Count(projectId, token string) (int64, error)
	GenerateToken(projectId, name string) (*domain.Token, error)
	DeleteToken(tokenId string) error
}
