package main

// https://leetcode.cn/problems/create-binary-tree-from-descriptions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 根据描述创建二叉树

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMp := map[int]*TreeNode{}
	hasParent := map[*TreeNode]bool{}
	for _, description := range descriptions {
		valP := description[0]
		valC := description[1]
		related := description[2]

		if nodeMp[valP] == nil {
			nodeMp[valP] = &TreeNode{Val: valP}
		}
		parent := nodeMp[valP]

		if nodeMp[valC] == nil {
			nodeMp[valC] = &TreeNode{Val: valC}
		}
		child := nodeMp[valC]
		hasParent[child] = true

		if related == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}
	}

	for _, node := range nodeMp {
		if !hasParent[node] {
			return node
		}
	}
	return nil
}
