package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * 1e9)
	timeout := time.After(5 * 1e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-timeout:
			fmt.Println("over")
			return
		default:
			fmt.Println(".")
			time.Sleep(1e8)
		}
	}

}
