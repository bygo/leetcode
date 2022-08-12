package main

// https://leetcode.cn/problems/average-of-levels-in-binary-tree/

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
		for _, node := range que[:queL] {
			sum += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
		numsAvg = append(numsAvg, float64(sum)/float64(queL))
		que = que[queL:]
	}
	return numsAvg
}
