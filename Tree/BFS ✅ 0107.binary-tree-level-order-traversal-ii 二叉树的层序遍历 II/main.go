package main

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

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
	var dep int
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		depsNums = append(depsNums, []int{})
		for _, q := range que[:queL] {
			if q.Left != nil {
				que = append(que, q.Left)
			}
			if q.Right != nil {
				que = append(que, q.Right)
			}
			depsNums[dep] = append(depsNums[dep], q.Val)
		}
		que = que[queL:]
		dep++
	}

	lo, hi := 0, len(depsNums)-1
	for lo < hi {
		depsNums[lo], depsNums[hi] = depsNums[hi], depsNums[lo]
		lo++
		hi--
	}

	return depsNums
}
