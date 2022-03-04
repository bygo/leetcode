package main

// https://leetcode-cn.com/problems/inorder-successor-in-bst-ii/

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// ❓ 二叉搜索树中的中序后继 II

func inorderSuccessor(node *Node) *Node {
	if node.Right == nil {
		for node.Parent != nil {
			if node.Parent.Left == node {
				// morris 最大后继
				// 如果root.Right为空，寻找以本树为左子树的父节点，找到返回，找不到返回nil
				// ⚠️ 相反，如果一直处于父节点的右子树，全部为false
				return node.Parent
			}
			node = node.Parent
		}
	} else { // 如果有root.Right，找最小值
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	return nil
}
