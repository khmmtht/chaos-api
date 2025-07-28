package main

import (
	"chaos-api/routes/api"
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	uri := os.Getenv("MONGODB_URI")
	client, _ := mongo.Connect(options.Client().ApplyURI(uri))
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	g := e.Group("api")
	api.AddChaosRoutes(g, client)
	api.AddProjectRoutes(g, client)
	api.AddGlobalRoutes(g)
	api.AddSimulateRoutes(g)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
