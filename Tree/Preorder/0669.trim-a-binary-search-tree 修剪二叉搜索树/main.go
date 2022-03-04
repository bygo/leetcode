package main

// https://leetcode-cn.com/problems/trim-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 修剪二叉搜索树

func trimBST(root *TreeNode, numL int, numR int) *TreeNode {
	if root == nil {
		return nil
	}
	if numR < root.Val { //太大 修剪，root 包括右子树全部砍掉
		return trimBST(root.Left, numL, numR)
	}
	if root.Val < numL { //太小 修剪，root 包括左子树全部砍掉
		return trimBST(root.Right, numL, numR)
	}
	root.Left = trimBST(root.Left, numL, numR)   //没问题的修剪左子树
	root.Right = trimBST(root.Right, numL, numR) //没问题的修剪右子树
	return root
}
