package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var level int
var counter int

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level = compute(root)

	if level == 0 {
		return 1
	}

	counter = int(math.Pow(2, float64(level)))
	left, right := 0, counter-1
	for left <= right {
		mid := (left + right) / 2
		if exist(mid, root) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return counter + left - 1
}

func exist(index int, node *TreeNode) bool {
	left, right := 0, counter-1
	l := level
	for 0 < l {
		l--
		mid := (left + right) / 2
		if index <= mid {
			right = mid - 1
			node = node.Left
		} else {
			left = mid + 1
			node = node.Right
		}
	}
	return node != nil
}

func compute(root *TreeNode) int {
	var l int
	for root.Left != nil {
		l++
		root = root.Left
	}
	return l
}
