// 生产者消费者问题

package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("put %v\n", i)
	}
}

func consumer(ch chan int) {
	for i := 0; i < 10; i++ {
		re := <-ch
		fmt.Printf("get %v\n", re)
	}
}

func pro_cosumer(size int) {
	ch := make(chan int, size)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(ch)
	}()

	go func() {
		defer wg.Done()
		consumer(ch)
	}()

	wg.Wait()
}

func main() {
	// 无缓冲: 生产一个消费一个
	pro_cosumer(0)

	// 带缓冲: 一次性生产多个
	//pro_cosumer(5)
}
