package main

import (
	"fmt"
	"time"
)

func a(msg string, sec int64) {
	fmt.Println(msg, "is ready")
	time.Sleep(time.Second)
}
func b() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic................")
		}
		fmt.Println("what happen")
	}()

	go a("tea", 3)
	go a("coffee", 2)

	go func() {
		go func() {
			defer func() {
				fmt.Println("get panic")
				return
			}()
			//panic("panicking")
		}()
	}()
}

type info struct {
	a int
	b string
}

func main() {
	fmt.Println("Antes...")
	go a("one a", 1)
	b()
	go a("two a", 1)
	time.Sleep(4 * 1e9)
	fmt.Println("Fim...")

	mp := make(map[string]interface{})
	mp["chenp"] = "googd"
	if mp["shit"] == nil {
		fmt.Println("shit is nil")
	}

	if mp["chenp"] == nil {
		fmt.Println("shit is nil")
	}

	i1 := new(info)
	if i1 == nil {
		fmt.Println("i1 is nil")
	}
}
