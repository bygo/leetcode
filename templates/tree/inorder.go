package tree

func inorder(root *TreeNode) []int {
	var res []int
	if root != nil {
		inorder(root.Left)
		res = append(res, root.Val) //中序输出
		inorder(root.Right)
	}

	return res
}
