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
}
