package gua

import (
	"fmt"
	"proj/utils"
)

type Gua struct {
	Index     int    `json:"index"`    // 周文王 卦序
	Name      string `json:"name"`     // 卦名
	GroupName string `json:"group"`    // 内外卦组合名称 例如 水雷屯
	Symbols   string `json:"symbols"`  // 图标
	YYGroup   []int  `json:"yy_group"` // 阴阳排列
}

// 补码 反码 朴码

// ZongGua 综卦-又称反卦和覆卦，指将一卦反复（颠倒）过来所得到的卦
func (g *Gua) ZongGua() Gua {
	ryg := utils.ReverseInitArr(g.YYGroup)
	rygStr := utils.IntArrToString(ryg)
	return GetGuaByYYGroup(rygStr)
}

// CuoGua 错卦-把一个卦的各个爻求反（阳变成阴，阴变成阳）
func (g *Gua) CuoGua() Gua {
	ryg := utils.ReverseYinYangArr(g.YYGroup)
	rygStr := utils.IntArrToString(ryg)
	return GetGuaByYYGroup(rygStr)
}

// HuGua 互卦-把主卦的最上爻和最下爻去掉，剩下四爻；以这四爻的上三爻作上卦(外卦)，下三爻作下卦(内卦)，构成新的六爻卦
func (g *Gua) HuGua() Gua {
	yg := g.YYGroup
	if len(yg) < 6 {
		return Gua{}
	}
	centerG := yg[1:5]
	shangGua := centerG[:3]
	xiaGua := centerG[1:]
	huGuaGroup := append(shangGua, xiaGua...)
	rygStr := utils.IntArrToString(huGuaGroup)
	return GetGuaByYYGroup(rygStr)
}

func (g *Gua) String() {
	/*
		Index     int    `json:"index"`    // 周文王 卦序
		Name      string `json:"name"`     // 卦名
		GroupName string `json:"group"`    // 内外卦组合名称 例如 水雷屯
		Symbols   string `json:"symbols"`  // 图标
		YYGroup   []int  `json:"yy_group"` // 阴阳排列
	*/
	fmt.Println(g.Symbols, g.Name, g.GroupName, g.YYGroup)

}
