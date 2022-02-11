package main

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的层序遍历

func levelOrder(root *TreeNode) [][]int {
	var depsNums [][]int
	if root == nil {
		return depsNums
	}

	var dep int
	var que = []*TreeNode{root}
	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}
		depsNums = append(depsNums, []int{})
		for _, q := range que[:cnt] {
			depsNums[dep] = append(depsNums[dep], q.Val)
			if q.Left != nil {
				que = append(que, q.Left)
			}
			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[cnt:]
		dep++
	}

	return depsNums
}
