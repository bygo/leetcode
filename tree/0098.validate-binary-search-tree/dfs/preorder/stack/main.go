package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack = []*TreeNode{root}
	var minS = []int{-1 << 63}
	var maxS = []int{1<<63 - 1}
	for len(stack) > 0 {
		//pop
		pre := len(stack) - 1
		root, min, max := stack[pre], minS[pre], maxS[pre]
		stack, minS, maxS = stack[:pre], minS[:pre], maxS[:pre]

		//push
		for root != nil {
			if root.Val <= min || max <= root.Val {
				return false
			}
			minS = append(minS, root.Val)
			maxS = append(maxS, max)
			stack = append(stack, root.Right)
			max = root.Val
			root = root.Left
		}
	}
	return true
}
