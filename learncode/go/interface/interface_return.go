package main

import "fmt"

type AA struct {
	a int
}

func (a *AA) Add() {
	a.a++
}

func (a *AA) Print() {
	fmt.Println(*a)
}

func (a *AA) Get() interface{} {
	return a.a
}

type Adder interface {
	Add()
	Get() interface{}
}

type NewAdder struct {
	Adder
	i int
}

func func_return(adder Adder) Adder {
	return &NewAdder{adder, 5}
}

func main() {
	var a Adder
	a = &AA{15}

	re := func_return(a)
	re.Add()
	fmt.Println(re.Get())
	//re.Print()
	re.Print()
}
