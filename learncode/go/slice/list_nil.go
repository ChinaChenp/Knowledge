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

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := make([]int, 20)
	for i, v := range s1 {
		s2[i] = v
	}

	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 0)
	s4 := s3[:0]
	fmt.Println(s3, s4)

	fmt.Println(s1[:2])
}
