package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var data = make(map[string]ValueData)

type ValueData struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatTime time.Time `json:"creat_time"`
	LifeTime  time.Time `json:"life_time"`
}

func main() {
	go updateTimer()
	e := echo.New()
	e.GET("/api/value/:key", getValue)
	e.POST("/api/value/:key", setValue)
	e.DELETE("/api/value/:key", deleteValue)
	e.Logger.Fatal(e.Start(":1323"))
}

//查询键值 get
func getValue(c echo.Context) (err error) {
	var valueData ValueData
	var is bool
	Key := c.Param("key")
	valueData, is = data[Key]
	if is {
		return c.JSON(http.StatusOK, valueData)
	} else {
		return c.String(http.StatusOK, "not found")
	}
}

//设置键值 post 必须参数：value 可选参数：life_time 默认永久
func setValue(c echo.Context) (err error) {
	valueData := new(ValueData)
	valueData.Key = c.Param("key")
	_, is := data[valueData.Key]
	if is {
		return c.String(http.StatusOK, "already exists")
	}
	if err = c.Bind(valueData); err != nil {
		return
	}
	if valueData.Value == "" {
		return c.String(http.StatusOK, "parameters required")
	}
	valueData.CreatTime = time.Now()
	if !valueData.LifeTime.IsZero() && valueData.LifeTime.Before(valueData.CreatTime) {
		return c.String(http.StatusOK, "parameters error")
	}
	data[valueData.Key] = *valueData
	return c.String(http.StatusOK, "success")
}

//删除键值 delete
func deleteValue(c echo.Context) (err error) {
	key := c.Param("key")
	delete(data, key)
	return c.String(http.StatusOK, "success")
}

//定时移除过期键值
func updateTimer() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		for key := range data {
			if !data[key].LifeTime.IsZero() && time.Now().After(data[key].LifeTime) {
				delete(data, key)
			}
		}
	}
}
