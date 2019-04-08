package main

import "fmt"

func main() {
	var sa = make([]int, 0)
	for i := 0; i < 10; i++ {
		sa = append(sa, i)

	}

	// case 1， 目标切片没有分配空间，拷贝过后是空的
	var da = make([]int, 0, 10)
	var cc = 0
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\t%v\n", len(da), da)

	// case 2，拷贝的大小是dst和src中切片的最小值
	da = make([]int, 5)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

	// case 3，dst和src拷贝长度一样就是完整覆盖
	da = make([]int, 10)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)

	// case 4, dst长度长于src则只拷贝前面一部分
	da = make([]int, 15) // 默认是0
	da[10], da[11], da[len(da) - 1] = -1, -2, -10
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)
}
