package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int)    // send amount to deposit
var balances = make(chan int, 1) // receive balance
func Deposit(amount int)         { deposits <- amount }
func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			fmt.Println("write ok", balance)
		}
	}
}

func init() {
	go teller()
}

func main() {

	for i := 0; i < 100; i++ {
		a := 100
		go Deposit(a)
	}

	fmt.Println("go over")
	time.Sleep(2 * 1e9)
	fmt.Println("get reuslt = ", Balance())
	fmt.Println("get reuslt = ", Balance())
	time.Sleep(1 * 1e9)
}
