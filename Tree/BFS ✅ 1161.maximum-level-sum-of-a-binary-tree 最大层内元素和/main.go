package main

// https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最大层内元素和

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var que = []*TreeNode{root}
	var depMax, dep, sum int
	var max = -1 << 63
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		for _, q := range que[:queL] {
			sum += q.Val
			if q.Left != nil {
				que = append(que, q.Left)
			}

			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		if max < sum {
			max = sum
			depMax = dep
		}
		sum = 0
		que = que[queL:]
	}

	return depMax
}
