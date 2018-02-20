package main

import "fmt"

type List struct {
	a    int
	next *List
}

// TODO : 链表倒排
func (l *List) Insert(a int) {
	head := l
	for head.next != nil {
		head = head.next
	}

	head.next = &List{a, nil}
}

func (l *List) Print() {
	head := l.next
	for head != nil {
		fmt.Printf("%d ", head.a)
		head = head.next
	}
	fmt.Println()
}

func (l *List) Reserve() {
	head := l.next
	if head != nil {
		if head.next != nil {
			l.next.Reserve()
		}
		fmt.Printf("%d ", head.a)
	}
}

func main() {
	list := List{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Print()

	list.Reserve()
}
