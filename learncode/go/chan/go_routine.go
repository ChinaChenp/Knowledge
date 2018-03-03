package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for x := 0; x < 20; x++ {
			ch1 <- x
		}
		close(ch1)
	}()

	go func() {
		for {
			x := <-ch1
			ch2 <- x * x
		}
		close(ch2)
	}()

	for x := range ch2 {
		fmt.Println(x)
	}
}
