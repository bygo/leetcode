package tree

func postorder(root *TreeNode) []int {
	var res []int
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val) //后序输出
	}
	return res
}
