package main

import "fmt"

type Info struct {
	age  int
	name string
}

type Any interface{}

func main() {
	any := make([]Any, 0)
	any = append(any, 1)
	any = append(any, "string")
	any = append(any, 1.62)
	any = append(any, true)

	info := &Info{3, "chenping"}
	any = append(any, info)

	for _, t := range any {
		switch t.(type) {
		case int:
			fmt.Printf("Type int %T\n", t)
		case string:
			fmt.Printf("Type string %T\n", t)
		case bool:
			fmt.Printf("Type boolean %T\n", t)
		case *Info:
			fmt.Printf("Type pointer to Info %T\n", t)
		default:
			fmt.Printf("Unexpected type %T\n", t)
		}
	}

	var shit interface{}
	re := fmt.Sprintf("%v", shit)
	fmt.Printf("%v\n", re)
}
