package main

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的层序遍历 II

func levelOrderBottom(root *TreeNode) [][]int {
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
		depsNums[depCur] = append(depsNums[depCur], node.Val)
		dfs(node.Left)
		dfs(node.Right)
		depCur--
	}
	dfs(root)

	lo, hi := 0, len(depsNums)-1
	for lo < hi {
		depsNums[lo], depsNums[hi] = depsNums[hi], depsNums[lo]
		lo++
		hi--
	}
	return depsNums
}
