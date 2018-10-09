package main

import (
	"fmt"
	"unsafe"
)

type Info struct {
	Age *int
	Num int
}

type From struct {
	Name string
	Detail *Info
}

type To struct {
	Name string
	Detail *Info
}

func copyfrom() *To {
	age := 29
	src := &From{
		Name : "chenp",
		Detail : &Info{
			Age: &age,
			Num: age,
		},
	}

	return (*To)(unsafe.Pointer(src))
}

func main() {
	var a = 15
	for i := 0; i < 1; i++ {
		var age= i
		var from From
		from.Name = "chenp" + fmt.Sprintf("_%d", i)
		from.Detail = &Info{&age, age}
		//bb := 888888888
		bb := 777777777
		var to To
		to = To(from)
		from.Detail.Age = &a
		fmt.Printf("%#v, %v, %v\n", from, *from.Detail, *from.Detail.Age)
		fmt.Println("----------")
		fmt.Printf("%#v, %v, %v, %v\n", to, *to.Detail, *from.Detail.Age, bb)
	}

	re := copyfrom()
	fmt.Printf("%#v, %v, %v\n", re, *re.Detail, *re.Detail.Age)
}