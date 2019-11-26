package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func shuffleSlice(a []int) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

// 洗牌算法
func Shuffle(a []int) {
	rand.Seed(time.Now().Unix())
	for i := len(a) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffleSlice(a)
	fmt.Println(a)

	sort.Ints(a)
	fmt.Println(a)
	Shuffle(a)
	fmt.Println(a)
}
