package main

import (
	"fmt"
	"proj/config"
	"proj/utils"
	"time"
)

func main() {

	//config.BuildJson()
	gua64 := config.GetGua64()
	Memory(gua64)

}

func Memory(GuaArr []config.Gua) {
	for {
		luckyNum := utils.RndNum(0, len(GuaArr)-1)
		gua := GuaArr[luckyNum]
		fmt.Print("					", gua.Symbols)
		time.Sleep(time.Second * 3)
		group := gua.Group
		runes := []rune(gua.Group)
		if len(runes) >= 2 {
			group = string(runes[:2])
		}
		fmt.Print(" ", group)
		time.Sleep(time.Second * 2)
		fmt.Println(" ", gua.Name, gua.Index)
	}
}
