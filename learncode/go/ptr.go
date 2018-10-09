package main

import "fmt"

type Info5 struct {
	A int
	B string
}

func main() {
	//var ptr *Info5
	//fmt.Println(ptr.A)

	var a = 15
	var ptr1 = &a
	var in interface{}
	in = ptr1
	fmt.Printf("%+v\n", *in.(*int))
	*ptr1 = 16
	fmt.Printf("%+v", *in.(*int))
}
