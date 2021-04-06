package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int     //最大值
var res []int   //结果
var cur int     //当前
var counter int //当前计数

func findMode(root *TreeNode) []int {
	res, max, cur, counter = []int{}, 0, 0, 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		if root.Val != cur {
			counter = 0
		}
		counter++
		if max < counter {
			max = counter
			res = []int{root.Val}
		} else if max == counter {
			res = append(res, root.Val)
		}
		cur = root.Val
		dfs(root.Right)
	}
}
