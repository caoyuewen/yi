package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Gua struct {
	Index   int    `json:"index"`
	Name    string `json:"name"`
	Group   string `json:"group"`
	Symbols string `json:"symbols"`
}

func GetGua64() []Gua {
	file, err := os.ReadFile("config/config.json")
	if err != nil {
		panic(err.Error())
	}
	var Gua64 []Gua
	err = json.Unmarshal(file, &Gua64)
	if err != nil {
		panic(err)
	}
	return Gua64
}

// ================================ 原始构建 ===========================

// GuaArr 64卦
var GuaArr [64]Gua
var groups []string

func BuildJson() {
	initGuaGroup()
	initGuaArr()

	marshal, err := json.Marshal(GuaArr)
	if err != nil {
		panic(err)
	}
	writeFile("config/config.json", marshal)
}

func initGuaArr() {
	open, err := os.ReadFile("config/gua.txt")
	if err != nil {
		panic(err.Error())
	}
	guaArr := strings.Split(string(open), "\n")
	for i := 0; i < len(guaArr); i++ {
		split := strings.Split(guaArr[i], "\t")
		for j := 0; j < len(split); j++ {
			g := strings.Fields(split[j])
			if len(g) != 3 {
				panic("config is err")
			}
			index, _ := strconv.Atoi(g[0])
			gua := Gua{
				Name:    g[1],
				Symbols: g[2],
				Index:   index,
				Group:   getGroupByGua(g[1]),
			}
			GuaArr[index-1] = gua
		}
	}
}

func initGuaGroup() {
	open, err := os.ReadFile("config/group.txt")
	if err != nil {
		panic(err.Error())
	}
	groups = strings.Fields(string(open))

}

func getGroupByGua(gua string) string {
	l := len([]rune(gua))
	for i := 0; i < len(groups); i++ {
		a := []rune(groups[i])
		if len(a) <= l {
			return gua
		}
		if string(a[len(a)-l:]) == gua {
			return string(a)
		}
	}
	return gua
}

func writeFile(filepath string, text []byte) {
	err := os.WriteFile(filepath, text, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("File written successfully!")
}
