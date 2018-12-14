package main

import "fmt"

func do() (str []byte, err error) {
	if false {
		return str, nil
	}

	return str, err
}

func main() {
	str, err := do()
	fmt.Println(str, err, len(str))
}
