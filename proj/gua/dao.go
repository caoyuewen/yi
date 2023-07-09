package gua

import (
	"fmt"
	"proj/utils"
)

const yin = 0
const yang = 1

var count int

func Dao1(g Gua) {
	count++
	if count > 2 {
		return
	}
	yg := g.YYGroup
	var all [][]int

	for i := 0; i < len(yg); i++ {
		if yg[i] == yin {
			//fmt.Print("	六", strconv.Itoa(i+1), ":")
			yg[i] = yang
		} else {
			//fmt.Print("	九", strconv.Itoa(i+1), ":")
			yg[i] = yin
		}
		a := make([]int, len(yg))
		copy(a, yg)
		all = append(all, a)
	}

	for i := 0; i < len(all); i++ {
		n := all[i]
		toString := utils.IntArrToString(n)
		Dao(GetGuaByYYGroup(toString))
	}
	Dao1(g.CuoGua())

}

func Dao(g Gua) {
	yg := g.YYGroup
	fmt.Println("	", g.Symbols, g.GroupName)
	for i := len(yg) - 1; i >= 0; i-- {
		nyg := make([]int, len(yg))
		copy(nyg, yg)
		if yg[i] == yin {
			//fmt.Print("	六", strconv.Itoa(i+1), ":")
			nyg[i] = yang
		} else {
			//fmt.Print("	九", strconv.Itoa(i+1), ":")
			nyg[i] = yin
		}
		ygStr := utils.IntArrToString(nyg)
		ng := GetGuaByYYGroup(ygStr)
		fmt.Println("	", ng.Symbols, ng.GroupName)
		//Dao(ng, i+1)
	}
	fmt.Println()
}
