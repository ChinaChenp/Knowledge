package main

import "fmt"

type A struct {
	a int
	b string
}

func main() {
	var in interface{}
	a := &A{1, "chenp"}
	in = a
	fmt.Printf("%p, %p\n", a, &in)
	in = nil
	re := in.(*A)
	fmt.Println(re)
}
