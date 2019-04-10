package main

import (
	"fmt"
	"unsafe"
)

type Info struct {
	List []int
	Kv   map[int]string
}

type DetailFrom struct {
	A string
	B int
	C *Info
}

type DetailTo struct {
	A string
	B int
	C *Info
}

func convert(from *DetailFrom) *DetailTo {
	return (*DetailTo)(unsafe.Pointer(from))
}

func main() {
	from := &DetailFrom{
		A: "chenp1",
		B: 50,
		C: &Info{
			List: []int{1, 2, 3, 4, 5},
			Kv:   map[int]string{1: "11", 2: "22", 3: "33"},
		},
	}

	to := convert(from)
	fmt.Println(to)
	fmt.Println(to.C)

	from.C.List = append(from.C.List, 6)
	from.C.Kv[4] = "44"
	fmt.Println(to.C)
	fmt.Println(from.C)
}
