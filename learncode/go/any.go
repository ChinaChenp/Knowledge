package main

import "fmt"

func main() {
	var v1 interface{} = 1
	v2 := v1.(int)
	fmt.Printf("v1 %d, v2 %d\n", v1, v2)
}
