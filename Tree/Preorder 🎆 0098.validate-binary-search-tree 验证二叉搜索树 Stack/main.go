package bfs

// https://leetcode-cn.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 验证二叉搜索树

type Pair struct {
	node *TreeNode
	min  int
	max  int
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack = []*Pair{{
		node: root,
	}}
	var min, max = -1 << 63, 1<<63 - 1
	for 0 < len(stack) {
		for root != nil {
			if root.Val <= min || max <= root.Val {
				return false
			}
			stack = append(stack, &Pair{
				node: root.Right,
				min:  root.Val,
				max:  max,
			})
			max = root.Val
			root = root.Left
		}
		top := len(stack) - 1
		pair := stack[top]
		stack = stack[:top]

		root = pair.node
		min = pair.min
		max = pair.max
	}
	return true
}
