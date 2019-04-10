package main

import (
	"fmt"
	"sort"
)

func PrintSlice(s []int) {
	for i, v := range s {
		fmt.Printf("index %d, value %d\n", i, v)
	}
	fmt.Printf("------------------------\n")
}

func returnSlice() []byte {
	s := []byte{'1', '2', '3', '4'}
	return s[0:2]
}

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("len s1 %d\n", len(s1))
	fmt.Printf("cap s1 %d\n", cap(s1))

	s1[0] = 0
	s1[1] = 1
	s1[2] = 2
	s1[3] = 3
	s1[4] = 4
	PrintSlice(s1)
	s2 := []int{11, 22, 33, 44}
	fmt.Printf("len s2 %d\n", len(s2))
	fmt.Printf("cap s2 %d\n", cap(s2))
	s2 = append(s2, 5, 6, 7, 8)
	PrintSlice(s2)
	PrintSlice(s2[:3])
	PrintSlice(s2[3:])

	s3 := []int{1, 2, 3, 4}
	s4 := []int{5, 6, 7}
	copy(s3, s4)
	PrintSlice(s3)
	copy(s4, s3)
	PrintSlice(s4)

	strslice := []byte{'1', '2', '3'}
	fmt.Printf("%v\n", string(strslice))

	re := returnSlice()
	fmt.Printf("----> %v, len %v, cap %v\n", re, len(re), cap(re))
	re1 := re[0 : 2 : len(re)+2]
	fmt.Printf("----> %v, len %v, cap %v\n", re1, len(re1), cap(re1))

	strslice1 := []string{"1", "2", "3", "4"}
	index := 3
	ss := append(strslice1[:index], strslice1[index+1:]...)
	fmt.Println(ss)

	strs := []string{"あ", "か", "さ", "た", "な", "は", "ま", "や", "ら", "わ"}
	fmt.Println("1", strs)
	sort.Strings(strs)
	fmt.Println("2", strs)
}
