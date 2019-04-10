package main

import (
	"fmt"
	"os"
)

// todo : 字典序二分查找
func strBinarySearch(str []byte, b byte) int {
	length := len(str)
	beg, end := 0, length-1
	for beg <= end {
		mid := (end + beg) / 2
		if b == str[mid] {
			return mid
		} else if b > str[mid] {
			beg = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	b := os.Args[1]

	str := "abcdefghijklmnopqrstuvwxyz"
	re := strBinarySearch([]byte(str), b[0])
	if re != -1 {
		fmt.Printf("%d->%c\n", re, str[re])
	} else {
		fmt.Println("not find\n")
	}
}
