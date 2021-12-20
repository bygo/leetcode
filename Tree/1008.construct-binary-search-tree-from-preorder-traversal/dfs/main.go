package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	l := len(preorder)
	if l == 0 {
		return nil
	}
	for 1 < l && preorder[0] < preorder[l-1] {
		l--
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:l]),
		Right: bstFromPreorder(preorder[l:]),
	}
}
