package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	ticker := time.Tick(5 * time.Second)

	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println("timeout")
			case <-ticker:
				fmt.Println("ticker")
			}
		}
	}()

	go func() {
		for {
			time.Sleep(1e9)
			timeout <- true
		}
	}()

	time.Sleep(10 * time.Second)

}
