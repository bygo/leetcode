package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m int

func findSecondMinimumValue(root *TreeNode) int {
	m = root.Val
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return -1 // -1 代表分支找到叶节点，都没找到一个符合的
	}
	if m < root.Val { //有大数就剪枝返回，因为我们找到了！
		return root.Val
	}

	l := dfs(root.Left)
	r := dfs(root.Right)
	if l == -1 { // -1 代表分支找到叶节点，都没找到一个符合的
		return r
	}
	if r == -1 { // -1 代表分支找到叶节点，都没找到一个符合的
		return l
	}

	return min(l, r) //两边都找到，就用最小数
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
