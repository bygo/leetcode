package main

// https://leetcode-cn.com/problems/find-bottom-left-tree-value/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 找树左下角的值
// ⚠️ 最底层的最左

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
		for _, node := range que[:queL] {
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		que = que[queL:]
	}
	return num
}
