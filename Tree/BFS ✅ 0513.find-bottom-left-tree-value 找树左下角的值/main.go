package main

// https://leetcode-cn.com/problems/find-bottom-left-tree-value/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 找树左下角的值

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var num int

	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		num = que[0].Val
		for _, q := range que[:queL] {
			if q.Left != nil {
				que = append(que, q.Left)
			}
			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[queL:]
	}
	return num
}
