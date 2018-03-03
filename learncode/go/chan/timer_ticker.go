package main

import "time"
import "fmt"

func main() {
	timeChan := time.NewTimer(time.Second).C

	//tickChan := time.NewTicker(time.Millisecond * 400).C

	doneChan := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 2)
		doneChan <- true
	}()

	for {
		select {
		case <-timeChan:
			fmt.Println("Timer expired")
		case <-doneChan:
			fmt.Println("Done")
		}
	}
}
