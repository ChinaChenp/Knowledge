package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	start := time.Now().Unix()
	fibN := fib(n) // slow
	end := time.Now().Unix()
	fmt.Printf("\rFibonacci(%d) = %d, use time %d\n", n, fibN, end-start)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
