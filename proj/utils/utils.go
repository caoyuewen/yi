package utils

import (
	"math/rand"
	"strconv"
	"time"
)

var count int64

// RndNum 获取一个随机数
func RndNum(startNum, endNum int) int {
	count++
	if count >= 1<<8 {
		count = 0
	}
	rand.NewSource(time.Now().UnixNano() + count)
	rnd := rand.Intn(endNum - startNum)
	return rnd + startNum
}

func IntArrToString(a []int) string {
	if len(a) <= 0 {
		return ""
	}
	var res string
	for i := 0; i < len(a); i++ {
		res += strconv.Itoa(a[i])
	}
	return res
}
