package main

// https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉搜索树

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var que = []*TreeNode{root}
	var minQ = []int{-1 << 63}
	var maxQ = []int{1<<63 - 1}
	for 0 < len(que) {
		var queL = len(que)
		for i, node := range que[:queL] {
			min, max := minQ[i], maxQ[i]
			if node.Val <= min || max <= node.Val {
				return false
			}
			if node.Left != nil {
				minQ = append(minQ, min)
				maxQ = append(maxQ, node.Val)
				que = append(que, node.Left)
			}
			if node.Right != nil {
				minQ = append(minQ, node.Val)
				maxQ = append(maxQ, max)
				que = append(que, node.Right)
			}
		}
		que, minQ, maxQ = que[queL:], minQ[queL:], maxQ[queL:]
	}

	return true
}
