package main

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的层序遍历

func levelOrder(root *TreeNode) [][]int {
	depsNums := [][]int{}
	var depCur = -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		depCur++
		for len(depsNums) < depCur+1 {
			depsNums = append(depsNums, []int{})
		}
		dfs(node.Left)
		dfs(node.Right)
		depsNums[depCur] = append(depsNums[depCur], node.Val)
		depCur--
	}
	dfs(root)
	return depsNums
}
