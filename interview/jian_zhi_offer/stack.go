package main

import (
	"errors"
	"fmt"
)

const STACK_MAX_SIZE = 65536

type Stack struct {
	iter int
	data [STACK_MAX_SIZE]int
}

func (s *Stack) Push(value int) error {
	if s.iter >= STACK_MAX_SIZE {
		return errors.New("stack is full")
	}

	s.data[s.iter] = value
	s.iter++
	return nil
}

func (s *Stack) Pop() int {
	if s.iter-1 < 0 {
		return -1
	}

	re := s.data[s.iter-1]
	s.iter--
	return re
}

func (s *Stack) Top() int {
	if s.iter-1 < 0 {
		return -1
	}
	return s.data[s.iter-1]
}

func (s *Stack) Size() int {
	return s.iter
}

func (s *Stack) Print() {
	fmt.Printf("(%d)==> ", s.Size())
	for i := s.iter - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d", s.data[i])
		} else {
			fmt.Printf("%d->", s.data[i])
		}
	}
	fmt.Println()
}

// TODO : min函数栈
type MinStack struct {
	data Stack
	min  Stack
}

func (s *MinStack) Push(v int) error {
	err := s.data.Push(v)
	if err != nil {
		return err
	}

	if s.min.Size() == 0 || s.min.Top() > v {
		s.min.Push(v)
	} else {
		s.min.Push(s.min.Top())
	}
	return nil
}

func (s *MinStack) Pop() int {
	s.min.Pop()
	return s.data.Pop()
}

func (s *MinStack) Min() int {
	return s.min.Top()
}

func main() {
	//s := &Stack{}
	//s.Push(4)
	//s.Push(6)
	//s.Push(3)
	//s.Push(9)
	//s.Push(10)
	//s.Print()
	//
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())
	//fmt.Println(s.Pop())

	sMin := &MinStack{}
	sMin.Push(4)
	fmt.Println(sMin.Min())
	sMin.Push(3)
	fmt.Println(sMin.Min())
	sMin.Push(5)
	fmt.Println(sMin.Min())
	sMin.Push(2)
	fmt.Println(sMin.Min())
	sMin.Push(3)
	fmt.Println(sMin.Min())

	fmt.Println()

	sMin.Pop()
	fmt.Println(sMin.Min())
	sMin.Pop()
	fmt.Println(sMin.Min())
	sMin.Pop()
	fmt.Println(sMin.Min())
	sMin.Pop()
	fmt.Println(sMin.Min())
	sMin.Pop()
	fmt.Println(sMin.Min())
}
