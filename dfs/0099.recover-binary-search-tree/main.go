package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历 o(1)
var last *TreeNode
var wrong []*TreeNode

func recoverTree(root *TreeNode) {
	last = nil
	wrong = make([]*TreeNode, 0)
	inorder(root)
	wrong[0].Val, wrong[1].Val = wrong[1].Val, wrong[0].Val
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	if last != nil && root.Val <= last.Val {
		if len(wrong) == 0 {
			wrong = append(wrong, last)
			wrong = append(wrong, root)
		} else {
			wrong[1] = root
		}
	}
	last = root
	inorder(root.Right)
}
