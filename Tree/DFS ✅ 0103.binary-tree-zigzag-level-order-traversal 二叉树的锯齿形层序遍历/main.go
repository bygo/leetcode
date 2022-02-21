package main

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的锯齿形层序遍历

func zigzagLevelOrder(root *TreeNode) [][]int {
	depsNums := [][]int{}
	var depCur = -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		depCur++
		if len(depsNums) <= depCur {
			depsNums = append(depsNums, []int{})
		}
		depsNums[depCur] = append(depsNums[depCur], node.Val)
		dfs(node.Left)
		dfs(node.Right)
		depCur--
	}
	dfs(root)
	for dep, nums := range depsNums {
		if dep%2 == 1 {
			lo, hi := 0, len(nums)-1
			for lo < hi {
				nums[lo], nums[hi] = nums[hi], nums[lo]
				lo++
				hi--
			}
		}
	}
	return depsNums
}
