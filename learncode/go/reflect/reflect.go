package main

import (
	"fmt"
	"reflect"
)

func convert(from, to interface{}) {
	ty := reflect.TypeOf(from)
	vl := reflect.ValueOf(from)
	fmt.Println(ty, vl)
	//to1 := (ty)(vl)
	//ty1 := reflect.TypeOf(to).Elem()
	//vl1 := reflect.ValueOf(to.(ty)).Elem()
	//vl1.Set(vl)
	//fmt.Println(vl1.Interface())
	//to = (ty1)(from)
}

func main() {
	var (
		a int64 = 5
		b int   = 6
	)

	convert(&a, &b)
	fmt.Println(a, b)
}
