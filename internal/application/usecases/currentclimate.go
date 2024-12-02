package usecases

import (
	"fmt"

	"br.com.cleiton/current-climate/internal/interface/services"
)

type CurrentClimateInterface interface {
	GetCurrentClimate(cep int) error
}

type CurrentClimate struct {
	cepApi     services.CepApiInterface
	climateApi services.ClimaApiInterface
}

func NewCurrentClimateUsecase(cepApi services.CepApiInterface, climateApi services.ClimaApiInterface) *CurrentClimate {
	return &CurrentClimate{
		cepApi:     cepApi,
		climateApi: climateApi,
	}
}

func (c CurrentClimate) GetCurrentClimate(cep int) error {
	if cep <= 0 {
		return fmt.Errorf("error to read cep value")
	}

	cepResponse, err := c.cepApi.GetLocation(cep)
	if err != nil {
		return fmt.Errorf("error to get location by cep, %w", err)
	}

	climateResponse, err := c.climateApi.GetCurrentClimate(cepResponse.Locality)
	if err != nil {
		return fmt.Errorf("error to get current climate, %w", err)
	}

	return nil
}
