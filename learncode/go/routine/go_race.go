package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	wg sync.WaitGroup
)

func incCounter() {
	defer wg.Done()

	counter++
}

func readCounter() int {
	defer wg.Done()
	return counter
}


func main() {
	counter := 10
	wg.Add(counter + 1)

	//for i := 0; i < counter; i++ {
	//	go incCounter()
	//}
	go incCounter()

	for i := 0; i < counter; i++ {
		go func() {
			re := readCounter()
			fmt.Println(re)
		}()
	}

	wg.Wait()
}
