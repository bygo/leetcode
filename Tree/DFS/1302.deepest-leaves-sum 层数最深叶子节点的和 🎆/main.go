package main

// https://leetcode-cn.com/problems/deepest-leaves-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 层数最深叶子节点的和

func deepestLeavesSum(root *TreeNode) int {
	var max, sum int
	var dfs func(node *TreeNode, dep int)
	dfs = func(node *TreeNode, dep int) {
		dfs(node.Left, dep+1)
		dfs(node.Right, dep+1)
		if node == nil {
			return
		}
		if max < dep {
			max = dep
			sum = node.Val
		} else if max == dep {
			sum += node.Val
		}
	}
	dfs(root, 0)
	return sum
}
