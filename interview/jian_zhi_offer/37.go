package main

import "fmt"

type Node1List struct {
	a    int
	next *Node1List
}

// TODO : 两个链表第一个公共节点
func (l *Node1List) InsertValue(a int) {
	head := l
	for head.next != nil {
		head = head.next
	}

	head.next = &Node1List{a, nil}
}

func (l *Node1List) AddList(list *Node1List) {
	head := l

	for head.next != nil {
		head = head.next
	}

	head.next = list.next
}

func (l *Node1List) Len() int {
	head := l
	var len int
	for head.next != nil {
		len++
		head = head.next
	}
	return len
}

func (l *Node1List) Print() {
	head := l.next
	for head != nil {
		fmt.Printf("%d ", head.a)
		head = head.next
	}
	fmt.Println()
}

// todo ： 长的先走diff步，然后同时走遇到第一个相同节点即是
func FindFirstCommonNode(l1, l2 *Node1List) *Node1List {
	len1 := l1.Len()
	len2 := l2.Len()

	if len1 == 0 || len2 == 0 {
		return nil
	}

	head1, head2 := l1.next, l2.next
	var diff int
	if len1 > len2 {
		diff = len1 - len2

		for i := 0; i < diff; i++ {
			head1 = head1.next
		}
	} else {
		diff = len2 - len1
		for i := 0; i < diff; i++ {
			head2 = head2.next
		}
	}

	for head1 != nil && head2 != nil && head1 != head2 {
		head1 = head1.next
		head2 = head2.next
	}

	return head1
}

func main() {
	l1 := Node1List{}
	l1.InsertValue(3)
	l1.InsertValue(9)
	l1.InsertValue(12)
	l1.InsertValue(8)
	l1.InsertValue(15)
	l1.Print()

	l2 := Node1List{}
	l2.InsertValue(9)
	l2.InsertValue(24)
	l2.InsertValue(98)
	l2.InsertValue(8)
	l2.Print()

	l3 := Node1List{}
	l3.InsertValue(11)
	l3.InsertValue(7)
	l3.InsertValue(34)

	l1.AddList(&l3)
	l2.AddList(&l3)
	l1.Print()
	l2.Print()

	fmt.Println()

	re := FindFirstCommonNode(&l1, &l2)
	fmt.Printf("first node is = %d\n", re.a)
}
