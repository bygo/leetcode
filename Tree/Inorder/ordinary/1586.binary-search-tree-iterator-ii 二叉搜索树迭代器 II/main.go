package main

// https://leetcode-cn.com/problems/binary-search-tree-iterator-ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉搜索树迭代器 II

type BSTIterator struct {
	idx   int
	nodes []*TreeNode
	stack []*TreeNode
	root  *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		nodes: []*TreeNode{},
		idx:   -1,
		root:  root,
	}
}

func (i *BSTIterator) HasNext() bool {
	return i.idx < len(i.nodes)-1 || i.root != nil
}

func (i *BSTIterator) Next() int {
	for i.root != nil {
		// 保存上下文 root包含 right 与 root
		i.stack = append(i.stack, i.root)
		i.root = i.root.Left
	}

	// 恢复上下文
	// 出栈
	top := len(i.stack) - 1
	for 0 <= top && i.root == nil {
		i.nodes = append(i.nodes, i.stack[top])
		i.root = i.stack[top].Right
		top--
	}
	i.stack = i.stack[:top+1]
	i.idx++
	return i.nodes[i.idx].Val
}

func (i *BSTIterator) HasPrev() bool {
	return 0 < i.idx
}

func (i *BSTIterator) Prev() int {
	i.idx--
	return i.nodes[i.idx].Val
}
