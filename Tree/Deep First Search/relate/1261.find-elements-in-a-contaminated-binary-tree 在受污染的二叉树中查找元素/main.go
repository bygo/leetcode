package main

// https://leetcode-cn.com/problems/find-elements-in-a-contaminated-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 在受污染的二叉树中查找元素

type FindElements struct {
	numMpBool map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	if root == nil {
		return FindElements{numMpBool: map[int]bool{}}
	}
	numMpBool := map[int]bool{0: true}
	var dfs func(node *TreeNode, parent int)
	dfs = func(node *TreeNode, parent int) {
		if node == nil {
			return
		}
		numCur := parent*2 + 1

		if node.Left != nil {
			numMpBool[numCur] = true
			dfs(node.Left, numCur)
		}
		if node.Right != nil {
			numCur += 1
			numMpBool[numCur] = true
			dfs(node.Right, numCur)
		}
	}
	dfs(root, 0)
	return FindElements{numMpBool: numMpBool}
}

func (f *FindElements) Find(target int) bool {
	return f.numMpBool[target]
}
