package main

import (
	"br.com.cleiton/current-climate/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	urlViaCep := viper.GetString("URL_VIA_CEP")
	urlClimate := viper.GetString("URL_CLIMATE")
	keyClimate := viper.GetString("KEY_CLIMATE")

	logrus.WithField("urlViaCep", urlViaCep).WithField("UrlClimate", urlClimate).WithField("key climate", keyClimate).Info("configs")

	server := internal.NewServer(urlClimate, keyClimate, urlViaCep)
	server.StartServer()
}
