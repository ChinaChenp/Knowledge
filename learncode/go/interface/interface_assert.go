package main

import "fmt"

type AA struct {
	a int
}

func (a *AA) Add() {
	a.a++
}

func (a *AA) Get() interface{} {
	return a.a
}

type BB struct {
	b string
}

func (b *BB) Add() {
	b.b += "b"
}

func (b *BB) Addn(n int) {
	for i := 0; i < n; i++ {
		b.b += "b"
	}
}

func (b *BB) Get() interface{} {
	return b.b
}

type Adder interface {
	Add()
	Get() interface{}
}

func main() {
	var in interface{}
	s := "chenping"
	in = s
	fmt.Println(in)

	if s1, ok := in.(string); ok {
		fmt.Println("yes", s1)
	} else {
		fmt.Println("no")
	}
	//fmt.Println(s1)

	// interface pass
	adds := []Adder{&AA{0}, &BB{"b"}}
	for _, v := range adds {
		func(add Adder) {
			switch add.(type) {
			case *AA:
				add.(*AA).Add() // or add.Add()
				fmt.Println(add.Get())
			case *BB:
				add.(*BB).Add()
				add.(*BB).Addn(5)
				fmt.Println(add.Get())
			}
		}(v)
	}
}
