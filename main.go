package main

import (
	"tugasapi/view"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/weather", view.WeatherForecastHandler)

	port := ":8080" // Pilih port yang ingin Anda gunakan
	e.Start(port)
}
