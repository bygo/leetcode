package tree

func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)  //左节点 遍历
		println(root.Val)   //中序遍历
		inorder(root.Right) //右节点 遍历
	}
}
