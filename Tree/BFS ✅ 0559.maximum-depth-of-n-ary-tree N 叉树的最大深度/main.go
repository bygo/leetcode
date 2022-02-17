package main

// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/submissions/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的最大深度

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var dep int
	var que = []*Node{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		for _, node := range que[:queL] {
			for _, child := range node.Children {
				que = append(que, child)
			}
		}
		que = que[queL:]
	}
	return dep
}
