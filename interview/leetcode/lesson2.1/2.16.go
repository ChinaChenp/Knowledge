package main

import "fmt"

//You are given an n × n 2D matrix representing an image. Rotate the image by 90 degrees (clockwise).
//Follow up: Could you do this in-place?

// 思路：n x n 数组旋转90度，先按对角线翻转，然后按水平中轴线翻转

func swap(a, b *int) {
	*a, *b = *b, *a
}

func rotateImage(arr [][]int) [][]int {
	n := len(arr)

	// 沿对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ { // 只循环到对角线位置
			//fmt.Println(i, j)
			swap(&arr[i][j], &arr[n-1-j][n-1-i])
		}
	}

	// 沿水平中线翻转
	for i := 0; i < n/2; i++ { // 只遍历一半
		for j := 0; j < n; j++ {
			swap(&arr[i][j], &arr[n-1-i][j])
		}
	}

	return arr
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{3, 4, 6},
		{9, 2, 1},
	}

	re := rotateImage(arr)
	for _, v := range re {
		fmt.Println(v)
	}
}
