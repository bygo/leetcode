package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(node *TreeNode) bool {
	return node == nil || isBalanced(node.Left) &&
		abs(height(node.Left)-height(node.Right)) < 2 && //两边最大深度差  大于2
		isBalanced(node.Right)
}

//计算节点最大深度
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// 计算左右树深度->对比深度
