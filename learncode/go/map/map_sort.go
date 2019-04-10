package main

import (
	"fmt"
	"sort"
)

func main() {
	var keys []int

	m := map[string]int{"two": 2, "three": 3, "four": 4, "one": 1}
	for k, v := range m {
		keys = append(keys, v)
		fmt.Println(k, v)
	}

	sort.Ints(keys)
	fmt.Println(keys)
}
