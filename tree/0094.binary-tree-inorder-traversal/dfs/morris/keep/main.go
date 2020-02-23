package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var pre *TreeNode

	for root != nil{
		if root.Left == nil{
			res = append(res,root.Val)
			root = root.Right
		}else{
			pre = root.Left
			for pre.Right != nil && pre.Right != root{
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = root
				root = root.Left
			}else{
				root = pre.Right
				pre.Right = nil
				res = append(res,root.Val)
				root = root.Right
			}
		}
	}
	return res
}


