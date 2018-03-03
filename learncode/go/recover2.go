package main

import (
	"fmt"
	"time"
)

func Since(t time.Time) {
	fmt.Printf("%v\n", time.Since(t))
}

func Since1() {
	fmt.Printf("shit\n")
}

func main() {
	t := time.Now()
	defer func() {
		Since(t)
	}()

	defer func() {
		Since1()
	}()
	time.Sleep(2 * time.Second)
	return
}
