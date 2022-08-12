package main

// https://leetcode.cn/problems/maximum-width-of-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树最大宽度

func widthOfBinaryTree(root *TreeNode) int {
	var numWidth int
	var numsDep = [][2]int{}
	var dep int
	var dfs func(node *TreeNode, idx int)
	dfs = func(node *TreeNode, idx int) {
		if node == nil {
			return
		}
		if len(numsDep) <= dep {
			numsDep = append(numsDep, [2]int{
				idx, idx,
			})
		}
		if numsDep[dep][1] < idx {
			numsDep[dep][1] = idx
		}

		dep++
		dfs(node.Left, idx<<1)
		dfs(node.Right, idx<<1+1)
		dep--
	}
	dfs(root, 0)
	for _, nums := range numsDep {
		numCur := nums[1] - nums[0] + 1
		if numWidth < numCur {
			numWidth = numCur
		}
	}
	return numWidth
}
