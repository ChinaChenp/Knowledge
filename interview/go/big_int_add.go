package main

import (
	"fmt"
	"strings"
)

func addBigNum(a, b string) string {
	if len(a) <= len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}

	needAdd := false
	out := ""
	for i := len(a) - 1; i >= 0; i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'

		sum := aNum + bNum
		if needAdd {
			sum++
		}

		if sum > 9 {
			needAdd = true
			sum -= 10
		} else {
			needAdd = false
		}

		out = string(sum+'0') + out
	}

	if needAdd {
		out = "1" + out
	}

	return out
}

func main() {
	re := addBigNum("12345", "9239")
	fmt.Println(re)
}
