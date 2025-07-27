package main

import (
	"chaos-api/routes/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("api")
	api.AddChaosRoutes(g)
	api.AddGlobalRoutes(g)
	api.AddSimulateRoutes(g)

	e.Logger.Fatal(e.Start(":1323"))
}
