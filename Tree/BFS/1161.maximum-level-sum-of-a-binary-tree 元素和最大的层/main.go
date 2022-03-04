package main

// https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 元素和最大的层

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var que = []*TreeNode{root}
	var depMax, dep, sum int
	var sumMax = -1 << 63
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		sum = 0
		dep++
		for _, node := range que[:queL] {
			sum += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}

			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		if sumMax < sum {
			sumMax = sum
			depMax = dep
		}
		que = que[queL:]
	}

	return depMax
}
