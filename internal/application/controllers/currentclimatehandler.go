package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CurrentClimateHandler struct{}

func NewCurrentClimateHandler() *CurrentClimateHandler {
	return &CurrentClimateHandler{}
}

func (h *CurrentClimateHandler) CurrentClimate(c echo.Context) {

	cep, err := strconv.Atoi(c.Param("cep"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid zipcode")
	}

}
