package main

import (
	"fmt"
	"proj/gua"
	"proj/utils"
	"time"
)

func main() {

	/*	g := gua.GetGuaByName("屯")
		gua.Dao(g)

		gua.Dao1(g)*/

	Memory()

}

func Memory() {
	GuaArr := gua.GetGua64()
	for {
		luckyNum := utils.RndNum(0, len(GuaArr)-1)
		gua := GuaArr[luckyNum]
		fmt.Print("					", gua.Symbols)
		time.Sleep(time.Second * 3)
		group := gua.GroupName
		runes := []rune(gua.GroupName)
		if len(runes) >= 2 {
			group = string(runes[:2])
		}
		fmt.Print(" ", group)
		time.Sleep(time.Second * 2)
		fmt.Println(" ", gua.Name, gua.Index)
	}
}
