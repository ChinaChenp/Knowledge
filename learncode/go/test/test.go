package main

import "fmt"
import "time"
import "runtime"

func test(cost int, id int) {
	fmt.Println("BEGIN test ", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("END test", id) //, code)
}
func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 100; i++ {
		go test(16, i)
	}
	fmt.Println("main done")
	time.Sleep(100000000000000)
}
