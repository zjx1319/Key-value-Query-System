package model

import (
	"Project1/util"
	"time"
)

type ValueData struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatTime util.Time `json:"creat_time"`
	LifeTime  util.Time `json:"life_time"`
}

var colValue = make(map[string]ValueData)

func initModelValue() {
	//TODO:初始化数据
	go ValueTimer()
}

func ValueTimer() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		for key := range colValue {
			if !time.Time(colValue[key].LifeTime).IsZero() && time.Now().After(time.Time(colValue[key].LifeTime)) {
				delete(colValue, key)
			}
		}
	}
}

func GetValue(key string) (valueData ValueData, is bool, err error) {
	valueData, is = colValue[key]
	return
}

func SetValue(valueData ValueData) (err error) {
	colValue[valueData.Key] = valueData
	return
}

func DeleteValue(key string) (err error) {
	delete(colValue, key)
	return
}

func IsValueExist(key string) (is bool, err error) {
	_, is = colValue[key]
	return
}
