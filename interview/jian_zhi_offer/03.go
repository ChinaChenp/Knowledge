package main

import (
	"fmt"
	"os"
	"strconv"
)

// TODO : 从一个每一列每一行二位数组中查找一个数

// TODO : 1）右上顶点查找 2）当前值大于key舍弃当前列 3）当前值小于key舍弃当前行
func Search(array *[5][4]int, rows, colums, key int) (int, int, bool) {
	row := 0
	colum := colums - 1

	for row < rows && colum >= 0 {
		now := array[row][colum]
		if now == key {
			return row, colum, true
		} else if now > key {
			fmt.Println(now, ">")
			colum--
		} else {
			fmt.Println(now, "<")
			row++
		}
	}
	return -1, -1, false
}

func main() {
	arr := [5][4]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
		{10, 14, 18, 19},
	}

	key, _ := strconv.Atoi(os.Args[1])
	a, b, ok := Search(&arr, 4, 4, key)
	fmt.Println(a, b, ok)
}
