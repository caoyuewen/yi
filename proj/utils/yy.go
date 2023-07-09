package utils

func ReverseInitArr(arr []int) []int {
	var res []int
	for i := len(arr) - 1; i >= 0; i-- {
		res = append(res, arr[i])
	}
	return res
}

func ReverseYinYangArr(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		res = append(res, ReverseYinYang(arr[i]))
	}
	return res
}

func ReverseYinYang(o int) int {
	if o == 0 {
		return 1
	}
	return 0
}
