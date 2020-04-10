package tree

func preorder(root *TreeNode) {
	if root != nil {
		println(root.Val)
		preorder(root.Left)  //左节点 遍历
		preorder(root.Right) //右节点 遍历
	}
}
