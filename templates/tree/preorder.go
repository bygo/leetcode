package tree

func preorder(root *TreeNode) {
	if root != nil {
		println(root.Val)    //前序遍历
		preorder(root.Left)  //左节点 遍历
		preorder(root.Right) //右节点 遍历
	}
}
