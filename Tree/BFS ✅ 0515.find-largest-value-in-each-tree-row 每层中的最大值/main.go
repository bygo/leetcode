package main

// https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 每层中的最大值

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var numsMax []int
	var que = []*TreeNode{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		var max = -1 << 63
		for _, node := range que[:queL] {
			if max < node.Val {
				max = node.Val
			}
			if node.Left != nil {
				que = append(que, node.Left)
			}

			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		numsMax = append(numsMax, max)
		que = que[queL:]
	}

	return numsMax
}
