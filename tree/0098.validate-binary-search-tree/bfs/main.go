package bfs

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 2,
		},
	},
	Right: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
		},
	},
}

func main() {
	isValidBST(tree)
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue = []*TreeNode{root}
	var minQueue []int
	var maxQueue []int
	if root.Left != nil {
		minQueue = append(minQueue, -1<<63)
		maxQueue = append(maxQueue, root.Val)
	}

	if root.Right != nil {
		minQueue = append(minQueue, root.Val)
		maxQueue = append(maxQueue, 1<<63-1)
	}

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root.Left != nil {
			if root.Left.Val <= minQueue[0] || maxQueue[0] <= root.Left.Val {
				return false
			}
			queue = append(queue, root.Left)
			minQueue = append(minQueue[1:], minQueue[0])
			maxQueue = append(maxQueue[1:], root.Left.Val)
		}

		if root.Right != nil {
			if root.Right.Val <= minQueue[0] || maxQueue[0] <= root.Right.Val {
				fmt.Printf("%+v %+v %+v", minQueue, maxQueue,root.Right.Val)
				return false
			}
			queue = append(queue, root.Right)
			minQueue = append(minQueue[1:], root.Right.Val)
			maxQueue = append(maxQueue[1:], maxQueue[0])
		}
	}

	return true
}
