package mocks

import (
	"br.com.cleiton/current-climate/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockClimateApi struct {
	mock.Mock
}

func (m *MockClimateApi) GetCurrentClimate(locaty string) (*entities.CurrentClimate, error) {
	args := m.Called(locaty)
	return args.Get(0).(*entities.CurrentClimate), args.Error(1)
}
