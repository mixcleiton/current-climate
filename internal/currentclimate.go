package internal

import (
	"br.com.cleiton/current-climate/internal/application/controllers"
	"br.com.cleiton/current-climate/internal/application/usecases"
	"br.com.cleiton/current-climate/internal/infrastructure/services/viacepapi"
	"br.com.cleiton/current-climate/internal/infrastructure/services/weatherapi"
	"github.com/labstack/echo"
)

type Config struct {
	urlViaCep  string
	urlClimate string
	keyClimate string
}

func NewServer(urlClimate, keyClimate, urlViaCep string) *Config {
	return &Config{
		urlViaCep:  urlViaCep,
		urlClimate: urlClimate,
		keyClimate: keyClimate,
	}
}

func (c *Config) StartServer() {

	climateApi := weatherapi.NewWeatherApi(c.urlClimate, c.keyClimate)
	viaCepApi := viacepapi.NewViaCepApi(c.urlViaCep)
	currentClimateUsecase := usecases.NewCurrentClimateUsecase(viaCepApi, climateApi)
	currentClimateHandler := controllers.NewCurrentClimateHandler(currentClimateUsecase)

	e := echo.New()

	e.GET("/:cep", currentClimateHandler.CurrentClimate)

	e.Logger.Fatal(e.Start(":8080"))
}
