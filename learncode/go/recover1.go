package main

import (
	"fmt"
	"runtime"
	"time"
)

func nextRecover() {
	if r := recover(); r != nil {
		fmt.Println("next recover ==>", r)
	}
	fmt.Println("next recover ok")
}

func doRecover() {
	if r := recover(); r != nil {
		fmt.Println("do recover ==>", r)
		return
	}
}

func dopainc() {
	panic("not good")
}

func doPainc() {
	dopainc()
}

//func logStack() string {
//	buf := make([]byte, 1<<16)
//	stackSize := runtime.Stack(buf, true)
//	return string(buf[0:stackSize])
//}

func CrashStack() []byte {
	buf := make([]byte, 20)
	n := runtime.Stack(buf, false)
	return buf[0:n]
}

func doprint() {
	for {
		time.Sleep(time.Millisecond)
	}
}

func main() {
	defer nextRecover()
	defer doRecover() // ok
	doPainc()
	fmt.Println("ok")
}
