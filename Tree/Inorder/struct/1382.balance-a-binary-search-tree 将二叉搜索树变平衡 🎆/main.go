package main

// https://leetcode.cn/problems/balance-a-binary-search-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 将二叉搜索树变平衡

func balanceBST(root *TreeNode) *TreeNode {
	var nums []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		node.Left = nil
		nums = append(nums, node)
		dfs(node.Right)
		node.Right = nil
	}

	var build func(nodes []*TreeNode) *TreeNode
	build = func(nodes []*TreeNode) *TreeNode {
		nL := len(nodes)
		if nL == 0 {
			return nil
		}

		mid := nL / 2
		node := nodes[mid]
		node.Left = build(nodes[:mid])
		node.Right = build(nodes[mid+1:])
		return node
	}

	dfs(root)
	return build(nums)

}
