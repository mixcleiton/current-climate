package internal

import (
	"br.com.cleiton/current-climate/internal/application/controllers"
	"br.com.cleiton/current-climate/internal/application/usecases"
	"br.com.cleiton/current-climate/internal/infrastructure/services/viacepapi"
	"br.com.cleiton/current-climate/internal/infrastructure/services/weatherapi"
	"github.com/labstack/echo"
)

func StartServer() {

	urlViaCEP := "https://viacep.com.br"
	urlClimate := "https://api.weatherapi.com"
	keyClimate := "1875c3e3f3284760bfb201107240212"

	climateApi := weatherapi.NewWeatherApi(urlClimate, keyClimate)
	viaCepApi := viacepapi.NewViaCepApi(urlViaCEP)
	currentClimateUsecase := usecases.NewCurrentClimateUsecase(viaCepApi, climateApi)
	currentClimateHandler := controllers.NewCurrentClimateHandler(currentClimateUsecase)

	e := echo.New()

	e.GET("/:cep", currentClimateHandler.CurrentClimate)

	e.Logger.Fatal(e.Start(":8080"))
}
