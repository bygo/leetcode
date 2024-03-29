package bfs

// https://leetcode.cn/problems/validate-binary-search-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉搜索树

func isValidBST(root *TreeNode) bool {
	return dfs(root, -1<<63, 1<<63-1)
}

func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if !dfs(root.Left, min, root.Val) {
		return false
	}
	if root.Val <= min || max <= root.Val {
		return false
	}
	return dfs(root.Right, root.Val, max)
}
