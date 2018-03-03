package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("fuck on")
	err2 := errors.New("fuck on")
	if err1.Error() == err2.Error() {
		fmt.Printf("ok")
	}
}
