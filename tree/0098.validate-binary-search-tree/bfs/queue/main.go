package main

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
	var minMap = map[*TreeNode]int{root: -1 << 63}
	var maxMap = map[*TreeNode]int{root: 1<<63 - 1}

	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
			minMap[root.Left] = minMap[root]
			maxMap[root.Left] = root.Val
			if root.Left.Val <= minMap[root.Left] || maxMap[root.Left] <= root.Left.Val {
				return false
			}
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
			minMap[root.Right] = root.Val
			maxMap[root.Right] = maxMap[root]
			if maxMap[root.Right] <= root.Right.Val || root.Right.Val <= minMap[root.Right] {
				return false
			}
		}
	}

	return true
}
