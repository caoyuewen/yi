package gua

import (
	"proj/utils"
	"sync"
)

var searchEngine struct {
	nameMap    sync.Map
	yyGroupMap sync.Map
}

func loadYYGroupMap() {
	guaArr := GetGua64()
	for i := 0; i < len(guaArr); i++ {
		g := guaArr[i]
		searchEngine.yyGroupMap.Store(utils.IntArrToString(g.YYGroup), g)
	}
}

func loadNameMap() {
	guaArr := GetGua64()
	for i := 0; i < len(guaArr); i++ {
		g := guaArr[i]
		searchEngine.nameMap.Store(g.Name, g)
	}
}

func GetGuaByName(name string) Gua {
	value, ok := searchEngine.nameMap.Load(name)
	if ok {
		return value.(Gua)
	}
	return Gua{}
}

func GetGuaByYYGroup(name string) Gua {
	value, ok := searchEngine.yyGroupMap.Load(name)
	if ok {
		return value.(Gua)
	}
	return Gua{}
}
