package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的层序遍历 II

func levelOrderBottom(root *TreeNode) [][]int {
	var depsNums [][]int
	if root == nil {
		return depsNums
	}

	var que = []*TreeNode{root}
	var dep = -1
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		depsNums = append(depsNums, []int{})
		for _, node := range que[:queL] {
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			depsNums[dep] = append(depsNums[dep], node.Val)
		}
		que = que[queL:]
	}

	lo, hi := 0, len(depsNums)-1
	for lo < hi {
		depsNums[lo], depsNums[hi] = depsNums[hi], depsNums[lo]
		lo++
		hi--
	}

	return depsNums
}
