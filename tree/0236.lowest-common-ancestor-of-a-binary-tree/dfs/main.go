package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// nil 叶节点
		// p || q   root 也可能是答案
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil { //  l为nil 就以r为结果
		return r
	}
	if r == nil { //  r为nil 就以l为结果
		return l
	}
	// l和r 同时有值
	// 必定是公共节点
	return root
}

func lowestCommonAncestor_2(root, p, q *TreeNode) *TreeNode {
	if root != nil && root != p && root != q {
		l := lowestCommonAncestor(root.Left, p, q)
		r := lowestCommonAncestor(root.Right, p, q)
		if l == nil {
			return r
		}
		if r == nil {
			return l
		}
	}
	return root
}
