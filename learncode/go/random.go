package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//例如，rand.Intn 返回一个随机的整数 n，0 <= n <= 100。
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(2), ",")
	}

	fmt.Printf("\n-----------------------------------\n")
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 5; i++ {
		seq := rand.Perm(10)
		fmt.Printf("random %v\n", seq)
	}
}
