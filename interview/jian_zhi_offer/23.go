package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Value int

	Left  *Tree
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func NewRandom(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = Insert(t, (1+v)*k)
	}
	return t
}

func Insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{v, nil, nil}
	}
	if v < t.Value {
		t.Left = Insert(t.Left, v)
	} else {
		t.Right = Insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func PreOrder(node *Tree) {
	if node != nil {
		fmt.Printf("%d, ", node.Value)
		PreOrder(node.Right)
	}
	fmt.Println()
}

func InOrder(node *Tree) {
	if node != nil {
		InOrder(node.Left)
		fmt.Printf("%d, ", node.Value)
		InOrder(node.Right)
	}
	fmt.Println()
}

func PostOrder(node *Tree) {
	if node != nil {
		PostOrder(node.Left)
		fmt.Printf("%d, ", node.Value)
	}
	fmt.Println()
}

func LevelOrder(node *Tree) {
	if node == nil {
		return
	}

	l := list.New()
	l.PushBack(node)

	for l.Len() >= 0 {
		data := l.Front()
		fmt.Println(data)
		l.Remove(data)

		if node.Left != nil {
			l.PushBack(node.Left)
		}

		if node.Right != nil {
			l.PushBack(node.Right)
		}
	}
}

func main() {
	t := NewRandom(4)

	PreOrder(t)
	InOrder(t)
	PostOrder(t)
}
