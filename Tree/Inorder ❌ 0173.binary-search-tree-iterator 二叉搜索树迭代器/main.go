package main

// https://leetcode-cn.com/problems/binary-search-tree-iterator/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树迭代器

type BSTIterator struct {
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root: root}
}

func (bst *BSTIterator) Next() int {
	node := bst.n(true)
	return node.Val
}

func (bst *BSTIterator) HasNext() bool {
	node := bst.n(false)
	return node != nil
}

func (bst *BSTIterator) n(move bool) *TreeNode {
	var max *TreeNode

	for bst.root != nil {
		if bst.root.Left == nil {
			node := bst.root
			if move {
				bst.root = bst.root.Right
			}
			return node
		} else {
			max = bst.root.Left
			for max.Right != nil {
				max = max.Right
			}

			max.Right = bst.root

			bst.root, bst.root.Left = bst.root.Left, nil
		}
	}
	return nil
}
