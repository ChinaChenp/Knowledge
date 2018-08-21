package main

import "fmt"

func main() {
	var mp map[int64]string

	if mp == nil {
		fmt.Println("mp is nil")
	}

	for k, v := range mp {
		fmt.Println("nil cant be range", k, v)
	}
}
