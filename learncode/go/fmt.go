package main

import (
	"fmt"
	"strconv"
)

func check(v interface{}) {
	re := fmt.Sprintf("%v", v)
	fmt.Printf("RE = %v\n", re)
}

func main() {
	cell := "0999999911"
	cellInt, err := strconv.ParseInt(cell, 10, 64)
	if err != nil {

	}
	fmt.Printf("%v\n", cellInt)

	shit := `{"needDriverBiz":1,"needDriverBasic":1}`
	check(shit)
}
