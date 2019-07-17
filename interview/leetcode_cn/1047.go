package main

import "fmt"

func removeDuplicates(S string) string {
	if len(S) <= 1 {
		return S
	}

	stack := make([]byte, 0)
	stack = append(stack, S[0])
	for i := 1; i < len(S); i++ {
		if len(stack) == 0 || S[i] != stack[len(stack)-1] {
			stack = append(stack, S[i])
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}

	re := ""
	for i := 0; i < len(stack); i++ {
		re += string(stack[i])
	}
	return re
}

func main() {
	s := "abbaca"
	re := removeDuplicates(s)
	fmt.Println(re)
}
