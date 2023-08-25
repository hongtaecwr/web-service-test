package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"web-service-test/config"
	"web-service-test/routes"
)

func main() {
	config.ReadConfig()

	e := echo.New()
	e.HideBanner = true
	routes.NewRouter(e)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		panic(err)
	}

}
