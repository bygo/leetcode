package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	println((0 - 2) / 2)
}

var res [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)

	l := len(res)
	end := l / 2
	for i := 0; i < end; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return res
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}
