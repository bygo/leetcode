package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return root1 == nil && root2 == nil || root1 != nil && root2 != nil && root1.Val == root2.Val && (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) || flipEquiv(root1.Left, root2.Right) && flipEquiv(root2.Left, root1.Right))
}
