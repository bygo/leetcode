package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
var max int
var h map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	res = []int{}
	max = 0
	h = map[int]int{}
	dfs(root, 0)
	for k, v := range h {
		if v == max {
			res = append(res, k)
		} else if max < v {
			res = []int{k}
			max = v
		}
	}
	return res
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum += dfs(root.Left, sum) + dfs(root.Right, sum) + root.Val
	h[sum]++
	return sum
}
