package main

import (
	"Project1/config"
	"Project1/model"
	"Project1/router"
	"Project1/util"
	"github.com/labstack/echo/middleware"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config.InitConfig()
	model.InitModel()
	util.InitUtil()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	apiGroup := e.Group("/api/v1")
	router.InitRouter(apiGroup)

	log.Fatal(e.Start(config.Config.App.Address))
}
