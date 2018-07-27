package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func doIt(ctx context.Context) {
	fmt.Printf("-----2 %v-----\n", ctx)
	go func() {
		c := context.Background()
		fmt.Printf("-----3 %v-----\n", c)
	}()
	fmt.Printf("-----4 %v-----\n", context.Background())
	time.Sleep(2 * time.Second)
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")

	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "value1")
	ctx = context.WithValue(ctx, "key2", "value2")
	fmt.Printf("-----1 %v-----\n", ctx)
	go func(ctx context.Context) {
		doIt(ctx)
	}(ctx)

	time.Sleep(5 * time.Second)
}
