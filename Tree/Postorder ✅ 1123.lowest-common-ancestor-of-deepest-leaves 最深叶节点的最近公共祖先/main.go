package main

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最深叶节点的最近公共祖先
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode, dep int) (*TreeNode, int)
	dfs = func(node *TreeNode, dep int) (*TreeNode, int) {
		if node == nil {
			return nil, dep
		}
		l, depLeft := dfs(node.Left, dep+1)
		r, depRight := dfs(node.Right, dep+1)

		if depLeft == depRight { // 左右一样深，返回当前节点 与 深度，
			return node, depLeft
		} else if depLeft < depRight { // 更深节点，的公共祖先 = 自己
			return r, depRight
		} else { // hiRight < hiLeft
			return l, depLeft
		}
	}
	common, _ := dfs(root, 0)
	return common
}

// 865. 具有所有最深节点的最小子树
// https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes/
