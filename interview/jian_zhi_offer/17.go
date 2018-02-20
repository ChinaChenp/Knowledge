package main

import (
	"fmt"
)

type NodeList struct {
	a    int
	next *NodeList
}

// TODO : 倒数第N个节点
func (l *NodeList) Insert(a int) {
	head := l
	for head.next != nil {
		head = head.next
	}

	head.next = &NodeList{a, nil}
}

func (l *NodeList) Print() {
	head := l.next
	for head != nil {
		fmt.Printf("%d ", head.a)
		head = head.next
	}
	fmt.Println()
}

func (l *NodeList) Merge(other *NodeList) {
	l1 := l.next
	l2 := other.next

	iter := &NodeList{}
	if l1 == nil && l2 == nil {
		l.next = nil
	} else if l1 == nil && l2 != nil {
		l.next = l2
	} else if l1 != nil && l2 == nil {
		l.next = l1
	} else {
		for l1 != nil || l2 != nil {
			if l1 != nil && l2 != nil && l1.a >= l2.a {
				iter.Insert(l2.a)
				l2 = l2.next
			} else if l1 != nil && l2 != nil && l1.a < l2.a {
				iter.Insert(l1.a)
				l1 = l1.next
			} else if l1 == nil && l2 != nil {
				iter.Insert(l2.a)
				l2 = l2.next
			} else if l1 != nil && l2 == nil {
				iter.Insert(l1.a)
				l1 = l1.next
			}
		}
		l.next = iter.next
	}
}

func main() {
	list1 := &NodeList{}
	list1.Insert(1)
	list1.Insert(3)
	list1.Insert(5)
	list1.Insert(9)
	list1.Insert(11)
	list1.Print()

	list2 := &NodeList{}
	list2.Insert(1)
	list2.Insert(2)
	list2.Insert(6)
	list2.Insert(8)
	list2.Insert(11)
	list2.Insert(15)
	list2.Insert(19)
	list2.Print()

	list1.Merge(list2)
	fmt.Print("fuck off\n")
	list1.Print()
}
