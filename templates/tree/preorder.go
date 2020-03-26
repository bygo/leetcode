package tree

func preorder(root *TreeNode) {
	if root != nil {
		println(root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
}
