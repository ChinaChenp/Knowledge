package main

import "fmt"

type NList struct {
	a    int
	next *NList
}

// TODO : 倒数第N个节点
func (l *NList) Insert(a int) {
	head := l
	for head.next != nil {
		head = head.next
	}

	head.next = &NList{a, nil}
}

func (l *NList) Print() {
	head := l.next
	for head != nil {
		fmt.Printf("%d ", head.a)
		head = head.next
	}
	fmt.Println()
}

func (l *NList) NToTail(n int) int {
	head := l.next
	if n == 0 || l == nil {
		return -1
	}

	fast := head
	slow := head
	count := 0
	for i := 0; i < n; i++ {
		count++
		if count <= n && fast != nil {
			fast = fast.next
		} else {
			return -1
		}
	}

	for fast != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow.a
}

func main() {
	list := NList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	list.Print()

	re := list.NToTail(5)
	fmt.Println(re)
}
