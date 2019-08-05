package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main(){
	var value int32
	fmt.Println("origin value:", value, &value)

	go entry("1", &value, 1)

	go entry("2", &value, 2)

	go func() {
		tmp := atomic.LoadInt32(&value)
		fmt.Println(&tmp)
	}()

	time.Sleep(2*time.Second)
}

func entry(name string, value *int32, new int32) {

	old := atomic.SwapInt32(value, new)

	fmt.Println("goroutine name:", name, ", swap, value:", old, "new value", new)
}
