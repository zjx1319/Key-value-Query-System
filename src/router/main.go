package router

import "github.com/labstack/echo"

func InitRouter(e *echo.Group) {
	valueGroup := e.Group("/value")
	initValueGroup(valueGroup)
}
