package main

import "fmt"

// A Tree is a binary tree with integer values.
type tree struct {
	Value int

	Left  *tree
	Right *tree
}

func insert(t *tree, v int) *tree {
	if t == nil {
		return &tree{v, nil, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func preOrder(node *tree) {
	if node != nil {
		fmt.Printf("%d, ", node.Value)

		preOrder(node.Left)
		preOrder(node.Right)
	}
	fmt.Println()
}

// todo : 判断是否是平衡二叉树
func isBlance(t *tree) bool {
	if t == nil {
		return true
	}

	left := depth(t.Left)
	right := depth(t.Right)
	if left > right+1 || right < left-1 {
		return false
	}

	return isBlance(t.Left) && isBlance(t.Right)
}

// todo : 二叉树的深度
func depth(t *tree) int {
	if t == nil {
		return 0
	}

	left := depth(t.Left)
	right := depth(t.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	t := insert(nil, 5)
	insert(t, 9)
	insert(t, 4)
	insert(t, 3)
	preOrder(t)

	re := depth(t)
	fmt.Println(re)
}
