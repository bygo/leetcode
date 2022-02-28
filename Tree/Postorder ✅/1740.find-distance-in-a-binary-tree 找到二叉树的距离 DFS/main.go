package main

// https://leetcode-cn.com/problems/find-distance-in-a-binary-tree/

// ❓ 找到二叉树的距离
// ⚠️ 距离 = 边的数量

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	// 公共父节点
	var common func(node *TreeNode) *TreeNode
	common = func(node *TreeNode) *TreeNode {
		if node == nil || node.Val == p || node.Val == q {
			return node
		}

		left := common(node.Left)
		right := common(node.Right)
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}

		return node
	}

	// 路径
	var dist int
	var path func(node *TreeNode, dep int)
	path = func(node *TreeNode, dep int) {
		if node != nil {
			if node.Val == p || node.Val == q {
				dist += dep
			}

			path(node.Left, dep+1)
			path(node.Right, dep+1)
		}
	}

	c := common(root)
	path(c, 0)
	return dist
}
