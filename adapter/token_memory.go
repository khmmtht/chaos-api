package adapter

import (
	_interface "chaos-api/adapter/interface"
	"chaos-api/domain"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type MemoryTokenAdapter struct {
	Store sync.Map
}

func NewMemoryTokenAdapter() _interface.TokenAdapter {
	return &MemoryTokenAdapter{
		Store: sync.Map{},
	}
}

func (a *MemoryTokenAdapter) Count(projectId, token string) (int64, error) {
	result, ok := a.Store.Load(projectId)
	if !ok {
		return 0, errors.New("key does not exist")
	}

	return int64(len(result.([]domain.Token))), nil
}

func (a *MemoryTokenAdapter) GetTokens() ([]domain.Token, error) {
	tokens := make([]domain.Token, 0)
	a.Store.Range(func(key, value interface{}) bool {
		tokens = append(tokens, value.(domain.Token))
		return true
	})

	return tokens, nil
}

func (a *MemoryTokenAdapter) GetTokensByProjectId(projectId string) ([]domain.Token, error) {
	result, ok := a.Store.Load(projectId)
	if !ok {
		return make([]domain.Token, 0), errors.New("key does not exist")
	}

	return result.([]domain.Token), nil
}

func (a *MemoryTokenAdapter) GenerateToken(projectId, name string) (*domain.Token, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	token, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	result := domain.Token{
		Id:        id.String(),
		ProjectId: projectId,
		Value:     token.String(),
		Name:      name,
	}
	tokens, err := a.GetTokensByProjectId(projectId)
	if err != nil {
		return nil, err
	}

	tokens = append(tokens, result)
	a.Store.Store(projectId, tokens)

	return &result, nil
}

func (a *MemoryTokenAdapter) DeleteToken(tokenId string) error {
	a.Store.Range(func(key, value interface{}) bool {
		key = key.(string)
		tokens := value.([]domain.Token)
		for i, token := range tokens {
			if token.Id == tokenId {
				tokens = append(tokens[:i], tokens[i+1:]...)
				a.Store.Store(key, tokens)
				break
			}
		}

		return true
	})

	return nil
}
