package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val { // 都在左树
		return lowestCommonAncestor(root.Left, p, q)
	}

	if root.Val < p.Val && root.Val < q.Val { // 都在右树
		return lowestCommonAncestor(root.Right, p, q)
	}

	// 左右各一个
	return root
}
