package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var data = make(map[string]string)

type DataJson struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	e := echo.New()
	e.GET("/api/value/:key", getValue)
	e.POST("/api/value/:key", setValue)
	e.DELETE("/api/value/:key", deleteValue)
	e.Logger.Fatal(e.Start(":1323"))
}

func getValue(c echo.Context) (err error) {
	var dataJson DataJson
	var is bool
	dataJson.Key = c.Param("key")
	dataJson.Value, is = data[dataJson.Key]
	if is {
		return c.JSON(http.StatusOK, dataJson)
	} else {
		return c.String(http.StatusOK, "not found")
	}
}

func setValue(c echo.Context) (err error) {
	dataJson := new(DataJson)
	dataJson.Key = c.Param("key")
	if err = c.Bind(dataJson); err != nil {
		return
	}
	_, is := data[dataJson.Key]
	if is {
		return c.String(http.StatusOK, "already exists")
	}
	data[dataJson.Key] = dataJson.Value
	return c.String(http.StatusOK, "success")
}

func deleteValue(c echo.Context) (err error) {
	key := c.Param("key")
	delete(data, key)
	return c.String(http.StatusOK, "success")
}
