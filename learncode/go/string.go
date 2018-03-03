package main

import (
	"fmt"
	"strings"
)

func noEmpty(strs []string) []string {
	res := strs[:0]
	for _, s := range strs {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}

func main() {
	str := "1234567890"
	fmt.Println(str[:3], str[3:])
	fmt.Println(str[len(str)-2:])

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)
	fmt.Println(arr[:3], arr[3:])

	strs := []string{"", "11", "22", "33", "", "55"}
	fmt.Println(strs)
	fmt.Println(noEmpty(strs))

	s := strings.Split("abc, abc", "")
	fmt.Println(s, len(s))

	s = strings.Split("", ",")
	fmt.Println(s, len(s))

	s = strings.Split("abc,abc", ",")
	fmt.Println(s, len(s))
}
