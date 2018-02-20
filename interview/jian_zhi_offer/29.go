package main

import "fmt"

// TODO : 找出超过一半的数
func CheckMoreThanHalf(arr []int) int {
	times := 0
	last := arr[0]
	for i := 0; i < len(arr); i++ {
		if times == 0 {
			last = arr[i]
			times = 0
		}

		if arr[i] == last {
			times++
		} else {
			times--
		}
	}

	return last
}

func main() {
	arr := []int{1, 2, 3, 6, 2, 2, 5, 4, 2, 2}
	re := CheckMoreThanHalf(arr)
	fmt.Println(re)
}
