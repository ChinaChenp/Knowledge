package main

import "fmt"

func main() {
	var sa = make([]string, 0)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))

	}
	var da = make([]string, 0, 10)
	var cc = 0
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\t%v\n", len(da), da)
	da = make([]string, 5)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)
	da = make([]string, 10)
	cc = copy(da, sa)
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n", len(da), cc, da)
}
