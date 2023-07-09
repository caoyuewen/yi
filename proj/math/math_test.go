package math

import (
	"fmt"
	"testing"
)

func TestPingFangHe(t *testing.T) {
	x := 100
	fmt.Println(PingFangHe(x))
	fmt.Println(PingFangHeGongShi(x))

	fmt.Println(JieShen(4))

	fmt.Println(GetPi(1))

}
