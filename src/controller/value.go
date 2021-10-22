package controller

import (
	"Project1/model"
	"Project1/util"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func ValueGet(c echo.Context) (err error) {
	key := c.Param("key")
	valueData, ok, err := model.GetValue(key)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return util.ErrorResponse(c, http.StatusOK, "not found")
	}
	return util.SuccessResponse(c, http.StatusOK, valueData)
}

func ValueSet(c echo.Context) (err error) {
	key := c.Param("key")
	if is, _ := model.IsValueExist(key); is {
		return util.ErrorResponse(c, http.StatusOK, "already exist")
	}
	var valueData model.ValueData
	err = c.Bind(&valueData)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if valueData.Value == "" {
		return util.ErrorResponse(c, http.StatusOK, "parameters required")
	}
	valueData.CreatTime = util.Time(time.Now())
	if !time.Time(valueData.LifeTime).IsZero() && time.Time(valueData.LifeTime).Before(time.Time(valueData.CreatTime)) {
		return util.ErrorResponse(c, http.StatusOK, "parameters error")
	}
	valueData.Key = key
	err = model.SetValue(valueData)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(c, http.StatusOK, nil)
}

func ValueDelete(c echo.Context) (err error) {
	key := c.Param("key")
	if is, _ := model.IsValueExist(key); !is {
		return util.ErrorResponse(c, http.StatusOK, "not found")
	}
	err = model.DeleteValue(key)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(c, http.StatusOK, nil)
}
