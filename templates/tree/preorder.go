package tree

func preorder(root *TreeNode) []int {
	var res []int
	if root != nil {
		res = append(res, root.Val) //前序输出
		preorder(root.Left)         //左节点 遍历
		preorder(root.Right)        //右节点 遍历
	}
	return res
}
