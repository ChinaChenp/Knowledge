package main

import "fmt"

func main()  {
	deferCall()
}

func deferCall()  {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
		fmt.Println("打印后")
	}()
	panic("触发异常")
}