package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffleSlice(a []int) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffleSlice(a)
	fmt.Println(a)
}
