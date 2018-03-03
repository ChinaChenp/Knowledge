package main

import "fmt"

// todo : 数组循环移位，只允许附件2个变量
func RightShift(arr []byte, n int) string {
	length := len(arr)
	n %= length
	for n > 0 {
		tmp := arr[length - 1]
		for j := length - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = tmp
		n--
	}
	return string(arr)
}

func main() {
	str := "237ctdsaf4"
	re := RightShift([]byte(str), 12)
	fmt.Println(str, re)
}
