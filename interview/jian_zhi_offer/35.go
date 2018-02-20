package main

import "fmt"

func FindFirstNotRepeatChar(str string) byte {
	var mapping [256]byte
	for i := 0; i < len(str); i++ {
		mapping[str[i]]++
	}

	for i := 0; i < len(str); i++ {
		if mapping[str[i]] == 1 {
			return str[i]
		}
	}

	return '0'
}

func main() {
	str := "eadaccdeffb"
	re := FindFirstNotRepeatChar(str)
	fmt.Printf("%c\n", re)
}
