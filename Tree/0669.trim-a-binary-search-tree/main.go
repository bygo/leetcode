package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root != nil {
		if R < root.Val { //太大 修剪，包括右子树全部砍掉
			return trimBST(root.Left, L, R)
		}
		if root.Val < L { //太小 修剪 ，包括左子树全部砍掉
			return trimBST(root.Right, L, R)
		}
		root.Left = trimBST(root.Left, L, R)   //没问题的修剪左子树
		root.Right = trimBST(root.Right, L, R) //没问题的修剪右子树
	}
	return root
}
