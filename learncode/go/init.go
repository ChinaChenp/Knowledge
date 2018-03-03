package main

import "fmt"

type Info struct {
	Name  string
	Age   int
	Score float64
	Bad   bool
}

func main() {
	v1 := &Info{"chenp", 28, 99, true}
	v2 := new(Info)

	fmt.Println(v1, v2)
}
