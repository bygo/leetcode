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
		cnt := len(que)
		if cnt == 0 {
			break
		}
		sum = 0
		for _, q := range que[:cnt] {
			sum += q.Val
			if q.Left != nil {
				que = append(que, q.Left)
			}

			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		que = que[cnt:]
	}
	return sum
}
