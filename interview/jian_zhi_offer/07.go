package main

import (
	"fmt"
	"math/rand"
)

// TODO : 快速排序法
func QuickSort(arr []int, beg, end int32) {
	left, right := beg, end

	index := beg + rand.Int31()%(end-beg)
	base := arr[index]
	for left < right {
		// TODO : 从右边找比基准小的数
		for arr[right] > base && right > beg {
			right--
		}
		// TODO : 从左边找比基准大的数
		for arr[left] < base && left < end {
			left++
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if left < end {
		QuickSort(arr, left, end)
	}

	if right > beg {
		QuickSort(arr, beg, right)
	}

}

func main() {
	arr := []int{5, 2, 1, 5, 3, 56, 9, 4, 11, 8, 17, 34, 22, 1, 0}
	length := int32(len(arr)) - 1
	QuickSort(arr, 0, length)
	fmt.Println(arr)
}
