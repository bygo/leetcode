package main

// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的层平均值

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var que = []*TreeNode{root}
	var numsAvg []float64
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		var sum int
		for _, q := range que[:queL] {
			sum += q.Val
			if q.Left != nil {
				que = append(que, q.Left)
			}
			if q.Right != nil {
				que = append(que, q.Right)
			}
		}
		numsAvg = append(numsAvg, float64(sum)/float64(queL))
		que = que[queL:]
	}
	return numsAvg
}
