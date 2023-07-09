package math

func PingFangHe(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i * i
	}
	return res
}

func PingFangHeGongShi(n int) int {
	return (2*n*n*n + 3*n*n + n) / 6
}

// JieShen !
func JieShen(n int) float64 {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return float64(res)
}

// QiuHe ∑
func QiuHe(n int, f func(k int) float64) float64 {
	var res float64
	for i := 1; i <= n; i++ {
		res += f(i)
	}
	return res
}

func GetPi(k int) float64 {

	return 0
}
