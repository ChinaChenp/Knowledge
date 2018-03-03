package main

import (
	"fmt"
	"sort"
)

func uniqueSlices(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if !unique[elem] {
			us = append(us, elem)
			unique[elem] = true
		}
	}

	return us
}

func MergeSlice(s1, s2 []string) []string {
	slice := make([]string, 0)
	for i := range s1 {
		slice = append(slice, s1[i])
	}

	for i := range s2 {
		slice = append(slice, s2[i])
	}

	sort.Strings(slice)
	return uniqueSlices(slice)
}

func Slice2Str(s []string, step string) string {
	var str string
	for i := range s {
		if i+1 != len(s) {
			str += s[i] + step
		} else {
			str += s[i]
		}
	}
	return str
}

func main() {
	s1 := []string{"", "5", "9", "3", "2"}
	//s2 := []string{"", "5", "91", "3", "2", "90"}
	var s2 []string
	re := MergeSlice(s1, s2)
	for _, v := range re {
		fmt.Printf("%s,", v)
	}
	fmt.Println()

	str := Slice2Str(re, "%")
	fmt.Println(str)
}
