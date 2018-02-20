package main

import (
	"container/list"
	"fmt"
)

// TODO : 圆圈中最后剩下的数字
func LastNumberInCricle(m, n int) int {
	list := list.New()
	for i := 0; i < m; i++ {
		list.PushBack(i)
		fmt.Printf("%d,", i)
	}
	fmt.Println()

	count := 1
	e := list.Front()
	for list.Len() > 1 {
		if count == n {
			count = 1
			fmt.Printf("remove %v\n", e.Value)
			list.Remove(e)
		}

		e = e.Next()
		if e == nil {
			count++
			e = list.Front()
		}

		count++
	}

	return e.Value.(int)
}

func main() {
	re := LastNumberInCricle(5, 3)
	fmt.Println(re)
}
