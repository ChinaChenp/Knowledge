package main

import "fmt"

func main ()  {
	l := []int64{}
	if l == nil {
		fmt.Println("nil", len(l))
	} else {
		fmt.Println("not nil", len(l))
	}

	var ll []int64
	if ll == nil {
		fmt.Println("ll is nil")
	}
	fmt.Println("normal", len(ll))

	ll1 := []int64{11, 22, 33, 44, 55}
	for i, v := range ll1 {
		fmt.Println(i, v)
	}

	ll2 := make([]int, 0)
	if ll2 == nil {
		fmt.Println("ll2 nil", len(ll2))
	}

	if len(ll2) == 0 {
		fmt.Println("ll2 len 0")
	}
}
