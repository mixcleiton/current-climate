package services

import "br.com.cleiton/current-climate/internal/domain/entities"

type CepApiInterface interface {
	GetLocation(cep int) (*entities.CEP, error)
}
