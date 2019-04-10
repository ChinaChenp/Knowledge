package main

import "fmt"

type Infoer interface {
	Print()
	Talk() (str string)
}

func main() {
	var info Infoer
	info.Talk()
	fmt.Println("shit")
}
