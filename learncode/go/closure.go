package main
import (
	"fmt"
)
func main() {
	// 闭包
	var j int = 5
	a := func() (func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()
	j *= 2
	a()

	// 匿名函数
	re := func(a, b int) int {
		return a + b
	}(5, 6)
	fmt.Printf("re = %d\n", re)

	f := func (b int) int {
		return b * 2
	}

	re = f(10)
	fmt.Printf("re = %d\n", re)
}
