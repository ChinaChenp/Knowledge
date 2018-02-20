package main

import (
	"fmt"
	"sort"
	"strconv"
)

// TODO ：把数组排成最小的数

type SortSlice []string

func (a SortSlice) Len() int {
	return len(a)
}
func (a SortSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a SortSlice) Less(i, j int) bool {
	return a[i]+a[j] < a[j]+a[i]
}

func FindMinNumber(arr []int) string {
	var strArry SortSlice
	for i := 0; i < len(arr); i++ {
		s := strconv.Itoa(arr[i])
		strArry = append(strArry, s)
	}

	sort.Sort(strArry)

	var re string
	for i := 0; i < len(strArry); i++ {
		//fmt.Printf("%s ", strArry[i])
		re += strArry[i]
	}
	return re
}

func main() {
	//arr := []int{4, 23, 723}
	//arr := []int{3, 32, 321}
	arr := []int{4, 35, 2}
	re := FindMinNumber(arr)
	fmt.Println(re)
}
