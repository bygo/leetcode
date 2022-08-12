package main

// https://leetcode.cn/problems/sum-of-nodes-with-even-valued-grandparent/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 祖父节点值为偶数的节点和

func sumEvenGrandparent(root *TreeNode) int {
	var sum int
	var dfs func(node *TreeNode, parent, grandparent bool)
	dfs = func(node *TreeNode, parent, grandparent bool) {
		if node == nil {
			return
		}
		if grandparent {
			sum += node.Val
		}
		dfs(node.Left, node.Val%2 == 0, parent)
		dfs(node.Right, node.Val%2 == 0, parent)

	}
	dfs(root, false, false)
	return sum
}
