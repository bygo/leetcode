package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var p, l, r *TreeNode

	for root != nil {
		l = root.Left //下一次循环

		root.Left = r  //上一个右节点
		r = root.Right //下一次循环

		root.Right = p //反转
		p = root       //下一次循环
		root = l
	}
	return p
}
