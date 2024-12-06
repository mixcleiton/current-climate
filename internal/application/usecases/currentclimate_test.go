package usecases

import (
	"strconv"
	"testing"

	"br.com.cleiton/current-climate/internal/domain/entities"
	"br.com.cleiton/current-climate/internal/interface/services/mocks"
	"github.com/stretchr/testify/assert"
)

func TestValidateGetClimateSuccess(t *testing.T) {
	mockClimateApi := new(mocks.MockClimateApi)
	mockCepApi := new(mocks.MockCepApi)

	climateUsecase := NewCurrentClimateUsecase(mockCepApi, mockClimateApi)

	cepFake := 999999

	mockCepApiResponse := &entities.CEP{
		Locality:       "CITY_TEST",
		Identification: strconv.Itoa(cepFake),
	}

	mockClimateApiResponse := &entities.CurrentClimate{
		Location: strconv.Itoa(cepFake),
		TempC:    1,
		TempF:    2,
		TempK:    0,
	}

	mockCepApi.On("GetLocation", cepFake).Return(mockCepApiResponse, nil)

	mockClimateApi.On("GetCurrentClimate", mockCepApiResponse.Locality).Return(mockClimateApiResponse, nil)

	currentClimateResponse, err := climateUsecase.GetCurrentClimate(cepFake)

	currentClimateSuccess := &entities.CurrentClimate{
		Location: "999999",
		TempC:    1,
		TempF:    2,
		TempK:    1 + valueConvertFahrenheit,
	}

	assert.NoError(t, err, nil)
	assert.Equal(t, currentClimateSuccess, currentClimateResponse)
}

func TestValidateGetClimateErrorApiCep(t *testing.T) {
	mockClimateApi := new(mocks.MockClimateApi)
	mockCepApi := new(mocks.MockCepApi)

	climateUsecase := NewCurrentClimateUsecase(mockCepApi, mockClimateApi)

	cepFake := 999999

	mockCepApi.On("GetLocation", cepFake).Return(&entities.CEP{}, ErrCep)

	_, err := climateUsecase.GetCurrentClimate(cepFake)

	assert.Error(t, err, ErrCep)
}

func TestValidateGetClimateErrorApiClimate(t *testing.T) {
	mockClimateApi := new(mocks.MockClimateApi)
	mockCepApi := new(mocks.MockCepApi)

	climateUsecase := NewCurrentClimateUsecase(mockCepApi, mockClimateApi)

	cepFake := 999999

	mockCepApiResponse := &entities.CEP{
		Locality:       "CITY_TEST",
		Identification: strconv.Itoa(cepFake),
	}

	mockCepApi.On("GetLocation", cepFake).Return(mockCepApiResponse, nil)

	mockClimateApi.On("GetCurrentClimate", mockCepApiResponse.Locality).Return(&entities.CurrentClimate{}, ErrClimate)

	_, err := climateUsecase.GetCurrentClimate(cepFake)

	assert.Error(t, err, ErrClimate)
}
