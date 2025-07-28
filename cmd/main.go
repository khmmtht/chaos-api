package main

import (
	"chaos-api/routes/api"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("api")
	api.AddChaosRoutes(g)
	api.AddGlobalRoutes(g)
	api.AddSimulateRoutes(g)

	e.Logger.Fatal(e.Start(":1323"))
}
