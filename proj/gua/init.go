package gua

import (
	"encoding/json"
	"os"
)

var gua64 []Gua

func init() {
	initGua64()
	loadYYGroupMap()
	loadNameMap()
}

func initGua64() {
	file, err := os.ReadFile("config/config.json")
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(file, &gua64)
	if err != nil {
		panic(err)
	}
}

func GetGua64() []Gua {
	res := make([]Gua, len(gua64))
	copy(res, gua64)
	return res
}
