package main

import "fmt"

func main() {
	slice := make([]string, 3, 11)

	var t trace
	t.Example(slice, "hello", 10)

	example(true, 'a', false)
}

type trace struct{}

func (t *trace) Example(slice []string, str string, i int) {
	fmt.Printf("receiver address: %p\n", t)
	//panic("Want stack trace")
}

func example(falg bool, c byte, f bool) {
	panic("package args")
}
