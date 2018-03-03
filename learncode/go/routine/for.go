package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)
	go func() {
		for {
			fmt.Print(1)
		}
	}()

	go func() {
		for {
			fmt.Print(2)
		}
	}()

	for {
		fmt.Print(0)
	}
}
