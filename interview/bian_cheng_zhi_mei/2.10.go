package main

import "fmt"

func FindMaxMin(arr []int) (int, int) {
	max, min := arr[0], arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] >= max {
			max = arr[i]
		}

		if arr[i] <= min {
			min = arr[i]
		}
	}
	return min, max
}

func main() {
	arr := []int {3, 6, 7, 9, 6, 5, 2, 3, 10, 2}
	max, min := FindMaxMin(arr)
	fmt.Println(max, min)
}
