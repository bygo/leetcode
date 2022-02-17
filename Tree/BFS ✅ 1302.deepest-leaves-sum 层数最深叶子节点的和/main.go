package main

// https://leetcode-cn.com/problems/deepest-leaves-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 层数最深叶子节点的和

func deepestLeavesSum(root *TreeNode) int {
	var sum int
	if root == nil {
		return sum
	}
	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		sum = 0
		for _, node := range que[:queL] {
			sum += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}

			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		que = que[queL:]
	}
	return sum
}
