package main

import (
	"fmt"
	"strings"
)

// todo : 字符串移位包含, 如给定s1=AABCD，s2=CDAA，返回true

func rightShift(str []byte) string {
	length := len(str)
	tmp := str[length-1]
	for j := length - 1; j > 0; j-- {
		str[j] = str[j-1]
	}
	str[0] = tmp

	return string(str)
}

func StrContant(s1, s2 string) bool {
	re := s1
	for i := 1; i < len(s1); i++ {
		re = rightShift([]byte(re))
		if strings.Contains(re, s2) {
			return true
		}
	}
	return false
}

func main() {
	str1 := "fjdia23rsdc"
	str2 := "dcfjd"
	re := StrContant(str1, str2)
	fmt.Println(re)
}
