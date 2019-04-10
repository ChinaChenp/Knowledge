package main

import (
	"fmt"
	"sync"
)

func doWork(req []int) []int {
	var wg sync.WaitGroup
	wg.Add(len(req))

	result := make(chan int)
	for _, v := range req {
		go func(a int) {
			defer wg.Done()
			if a%2 == 0 {
				fmt.Println(a)
				return
			}
			result <- a + 100
		}(v)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	l := make([]int, 0)
	for i := range result {
		l = append(l, i)
	}
	return l
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	re := doWork(list)
	fmt.Println(re)
}
