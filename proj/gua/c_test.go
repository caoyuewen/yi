package gua

import (
	"testing"
)

func TestGua(t *testing.T) {
	g := GetGuaByName("鼎")
	g.String()

	cuogua := g.CuoGua()
	cuogua.String()

	zongGua := g.ZongGua()
	zongGua.String()

	huGua := g.HuGua()
	huGua.String()

}
