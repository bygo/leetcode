package tree

func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		println(root.Val)
		inorder(root.Right)
	}
}
