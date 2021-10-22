package router

import (
	"Project1/controller"
	"github.com/labstack/echo"
)

func initValueGroup(group *echo.Group) {
	group.GET("/:key", controller.ValueGet)
	group.POST("/:key", controller.ValueSet)
	group.DELETE("/:key", controller.ValueDelete)
}
