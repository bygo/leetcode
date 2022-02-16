package bfs

// https://leetcode-cn.com/problems/validate-binary-search-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉搜索树

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack = []*TreeNode{root}
	var minStack, maxStack = []int{-1 << 63}, []int{1<<63 - 1}
	var min, max int
	for 0 < len(stack) {
		top := len(stack) - 1
		root, min, max = stack[top], minStack[top], maxStack[top]
		stack, minStack, maxStack = stack[:top], minStack[:top], maxStack[:top]

		for root != nil {
			if root.Val <= min || max <= root.Val {
				return false
			}
			stack, minStack, maxStack = append(stack, root.Right), append(minStack, root.Val), append(maxStack, max)
			max = root.Val
			root = root.Left
		}
	}
	return true
}
