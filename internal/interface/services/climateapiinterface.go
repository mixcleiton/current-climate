package services

import "br.com.cleiton/current-climate/internal/domain/entities"

type ClimaApiInterface interface {
	GetCurrentClimate(locaty string) (*entities.CurrentClimate, error)
}
