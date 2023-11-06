package view

import (
	"net/http"
	"tugasapi/controller"

	"github.com/labstack/echo/v4"
)

func WeatherForecastHandler(c echo.Context) error {
	location := c.QueryParam("location")
	if location == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Location is required"})
	}

	// Memanggil controller.FetchWeatherData dengan apiKey
	apiKey := "cc0a364fa16d48d0941135837230311" // Ganti dengan kunci API Anda
	weather, err := controller.FetchWeatherData(apiKey, location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching weather data"})
	}

	return c.JSON(http.StatusOK, weather)
}
