package main

// https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最深叶节点的最近公共祖先

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (*TreeNode, int)
	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return nil, -1
		}
		l, depLeft := dfs(node.Left)
		r, depRight := dfs(node.Right)

		if depLeft == depRight { // 左右一样深，返回当前节点 与 深度，
			return node, depLeft + 1
		} else if depLeft < depRight { // 更深节点，的公共祖先 = 自己
			return r, depRight + 1
		} else { // depRight < depLeft
			return l, depLeft + 1
		}
	}
	common, _ := dfs(root)
	return common
}

// 865. 具有所有最深节点的最小子树
// https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/
