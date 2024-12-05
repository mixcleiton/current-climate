package usecases

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"br.com.cleiton/current-climate/internal/domain/entities"
	"br.com.cleiton/current-climate/internal/interface/services"
)

type CurrentClimateInterface interface {
	GetCurrentClimate(cep int) (*entities.CurrentClimate, error)
}

const valueConvertFahrenheit = 273

var (
	ErrCep     = errors.New("error to get location by cep")
	ErrClimate = errors.New("error to get current climate")
)

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

func (c CurrentClimate) GetCurrentClimate(cep int) (*entities.CurrentClimate, error) {
	if cep <= 0 {
		return nil, fmt.Errorf("error to read cep value")
	}

	cepResponse, err := c.cepApi.GetLocation(cep)
	if err != nil || cepResponse == nil {
		log.Printf("err %s", err)
		return nil, ErrCep
	}

	log.WithField("cepResponse", cepResponse).Info("cep api result")
	climateResponse, err := c.climateApi.GetCurrentClimate(cepResponse.Locality)
	if err != nil {
		log.Printf("error in api current climate, err: %s", err)
		return nil, ErrClimate
	}

	climateResponse.TempK = climateResponse.TempC + valueConvertFahrenheit

	return climateResponse, nil
}
