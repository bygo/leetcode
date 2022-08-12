package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal/

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

	var dep = -1
	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		depsNums = append(depsNums, []int{})
		for _, node := range que[:queL] {
			depsNums[dep] = append(depsNums[dep], node.Val)
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		que = que[queL:]
	}

	return depsNums
}
