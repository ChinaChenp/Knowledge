package main

import "fmt"
import "errors"

func useage() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	//panic("shit")
	return
	defer fmt.Println("defer 3")
}

func userage1() (err error) {
	defer func() {
		if err == nil {
			fmt.Println("err is nil", err)
		} else {
			fmt.Println("err not nil", err)
		}
	}()

	err = errors.New("shit")
	return err
}

func main() {
	useage()
	userage1()
}
