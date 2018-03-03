package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Find(arr SortSlice, x int64) (int, bool) {
	maxLen := len(arr.FirstSort)
	minLen := len(arr.SecondSort)

	fmt.Println(arr.FirstSort[0][0], arr.SecondSort[maxLen-1][1])
	if x < arr.FirstSort[0][0] || x > arr.SecondSort[maxLen-1][1] {
		return 0, false
	}

	max := sort.Search(maxLen, func(i int) bool { return arr.SecondSort[i][1] >= x })
	fmt.Println("index max", max)

	min := sort.Search(minLen, func(i int) bool { return arr.FirstSort[i][0] <= x })
	fmt.Println("index min", min)

	if min < minLen || max < maxLen {
		if max < maxLen {
			for j := max; j >= 0; j-- {
				fmt.Println("max iter beg ", arr.SecondSort[j][0], " end ", arr.SecondSort[j][1])
				if arr.SecondSort[j][0] <= x && arr.SecondSort[j][1] >= x {
					return j, true
				}
			}
		}

		if min < minLen {
			for j := min; j < minLen; j++ {
				fmt.Println("min iter beg ", arr.FirstSort[j][0], " end ", arr.FirstSort[j][1])
				if arr.FirstSort[j][0] <= x && arr.FirstSort[j][1] >= x {
					return j, true
				}
			}
		}
	}

	return 0, false

	//	var min int64 = 1
	//	var max int64 = 111999999992
	//
	//	j := 0
	//	if i >= l && min <= x && max >= x {
	//		for j = i - 1; j >= 0; j-- {
	//			fmt.Println("out iter beg ", a[j][0], " end ", a[j][1])
	//			if a[j][0] <= x && a[j][1] >= x {
	//				return j, true
	//			}
	//		}
	//	}
	//
	//	j = 0
	//	if i < l {
	//
	//		for j = i; j >= 0; j-- {
	//			fmt.Println("iter beg ", a[j][0], " end ", a[j][1])
	//			if a[j][0] <= x && a[j][1] >= x {
	//				return j, true
	//			}
	//		}
	//	}
	//	return 0, false
}

type SortSlice struct {
	FirstSort  [][2]int64
	SecondSort [][2]int64
}

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	//a := [][2]int64{
	//	[2]int64{1, 1},
	//	[2]int64{3, 3},
	//	[2]int64{5, 5},
	//	[2]int64{6, 30},
	//	[2]int64{6, 6},
	//	[2]int64{9, 9},
	//	[2]int64{10, 12},
	//	[2]int64{14, 15},
	//	[2]int64{19, 26},
	//	[2]int64{110000001, 119000001},
	//	[2]int64{11100000000, 11199999999},
	//	[2]int64{111000000002, 111999999992},
	//}
	a := SortSlice{
		[][2]int64{
			[2]int64{1, 1},
			[2]int64{3, 3},
			[2]int64{5, 5},
			[2]int64{6, 30},
			[2]int64{6, 6},
			[2]int64{9, 9},
			[2]int64{10, 12},
			[2]int64{14, 15},
			[2]int64{19, 26},
			[2]int64{110000001, 119000001},
			[2]int64{11100000000, 11199999999},
			[2]int64{111000000002, 111999999992},
		},
		[][2]int64{
			[2]int64{1, 1},
			[2]int64{3, 3},
			[2]int64{5, 5},
			[2]int64{6, 6},
			[2]int64{9, 9},
			[2]int64{10, 12},
			[2]int64{14, 15},
			[2]int64{19, 26},
			[2]int64{6, 30},
			[2]int64{110000001, 119000001},
			[2]int64{11100000000, 11199999999},
			[2]int64{111000000002, 111999999992},
		},
	}

	index, ok := Find(a, int64(x))
	if ok {
		fmt.Printf("found %d at index %d in\n", x, index)
	} else {
		fmt.Printf("%d not found in \n", x)
	}
}
