package main

import "fmt"

var (
	intChan1 = make(chan int, 1)
	intChan2 chan int
	channels = []chan int{intChan1, intChan2}
	numbers  = []int{1, 2, 3}
)

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("case 0")
	case getChan(1) <- getNumber(1):
		fmt.Println("case 1")
	default:
		fmt.Println("default")
	}
}

func getChan(i int) chan int {
	fmt.Println("getChan", i)
	return channels[i]
}

func getNumber(i int) int {
	fmt.Println("getNumber", i)
	return numbers[i]
}
