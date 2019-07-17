package main

import "fmt"

func main() {

	hash := [256]rune{}
	s := "fdjafjvcjxklvjk"
	for _, v := range s {
		//fmt.Println(v)
		hash[v] = 1
	}

	for i := 0; i < 256; i++ {
		if hash[i] == 1 {
			fmt.Println(i)
		}
	}

	fmt.Println()
	s1 := "nihao 世界真好"
	for _, v := range s1 {
		fmt.Println(string(v))
	}

	s2 := []rune("nihao 世界真好")
	for i := range s2 {
		fmt.Println(string(s2[i]))
	}
}