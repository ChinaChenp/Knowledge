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
}

func doRecover() {
	if r := recover(); r != nil {
		//nextRecover()
		fmt.Println("----------------start----------------")
		fmt.Println("happen painc :", r) //prints: recovered => <nil>
		stack := CrashStack()
		fmt.Println("stack len", len(stack), "stack ->", string(stack))
	}
	fmt.Println("----------------end----------------")
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
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("recovered:", r) // ok
	//	}
	//}()
	defer func() {
		doRecover() // ok
	}()
	//doRecover() // not ok

	for i := 0; i < 10; i++ {
		go doprint()
	}
	//time.Sleep(time.Millisecond)
	//defer doRecover() //panic is not recovered

	// recover()   //doesn't do anything
	doPainc()
	// recover()   //won't be executed :)
	//time.Sleep(2 * time.Millisecond)

	fmt.Println("ok")
}
