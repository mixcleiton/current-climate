package viacepapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"br.com.cleiton/current-climate/internal/domain/entities"
	"br.com.cleiton/current-climate/internal/interface/services"
)

type viacepapi struct {
	urlViaCep string
}

func NewViaCepApi(urlViaCep string) services.CepApiInterface {
	return &viacepapi{
		urlViaCep: urlViaCep,
	}
}

func (v *viacepapi) GetLocation(cep int) (*entities.CEP, error) {
	service := fmt.Sprintf("/ws/%d/json/", cep)
	resp, err := http.Get(v.urlViaCep + service)
	if err != nil {
		return nil, fmt.Errorf("error to get request, %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error to read the response, %w", err)
	}

	var viaCepResponse ViaCepResponse
	err = json.Unmarshal(body, &viaCepResponse)
	if err != nil {
		return nil, fmt.Errorf("error to convert json in response")
	}

	return &entities.CEP{
		Locality:       viaCepResponse.Localidade,
		Identification: viaCepResponse.CEP,
	}, nil
}
