package main

import (
	"fmt"
	"time"
)

type Counter struct {
	cout int
}

var mapChan = make(chan map[string]Counter, 1)

//var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.cout++
				elem["count1"] = Counter{1}
				//elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]Counter{
			"count": Counter{},
		}
		//countMap := map[string]int{
		//	"count": 1,
		//}

		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v, %v. [sender]\n", countMap["count"], countMap["count1"])
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}
