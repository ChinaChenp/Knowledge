// 测试字符串是不是共享地址

package main

import (
	"fmt"
	"unsafe"
)

func copy_str(s1 string) {
	p := (*struct {
		tab  uintptr
		data uintptr
	})(unsafe.Pointer(&s1))

	s1 = "chenping1"

	fmt.Printf("s1 %v, in %v, %v\n", s1, p, p.tab)
}

func main() {
	s := "chenping"

	p := (*struct {
		tab  uintptr
		data uintptr
	})(unsafe.Pointer(&s))

	fmt.Printf("s1 %v, out %v, %v\n", s, p, p.tab)
	copy_str(s)
	fmt.Printf("s1 %v, out %v, %v\n", s, p, p.tab)
	return
}
