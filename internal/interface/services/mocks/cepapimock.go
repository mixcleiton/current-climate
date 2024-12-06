package mocks

import (
	"br.com.cleiton/current-climate/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockCepApi struct {
	mock.Mock
}

func (m *MockCepApi) GetLocation(cep int) (*entities.CEP, error) {
	args := m.Called(cep)
	return args.Get(0).(*entities.CEP), args.Error(1)
}
