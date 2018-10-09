package main

import (
	"fmt"
	"unsafe"
)

type DeliverType int64

const (
	DeliverType_Didi DeliverType = 1
	DeliverType_Shop DeliverType = 2
)

func convert(a *int) *int64 {
	return (*int64)(unsafe.Pointer(a))
}

func main()  {
	var A DeliverType
	fmt.Println(A)

	a := 15
	b := &a
	re := convert(b)
	fmt.Println(*re)

	var c *int
	re = convert(c)
	fmt.Println(re)
}