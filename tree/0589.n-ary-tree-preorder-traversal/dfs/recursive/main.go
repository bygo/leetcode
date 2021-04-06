package main

type Node struct {
	Val      int
	Children []*Node
}

//https://leetcode-cn.com/problemset/all/
func preorder(root *Node) []int {
	var res []int
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int) {
	if root != nil {
		*res = append(*res, root.Val)
		for _, v := range root.Children {
			dfs(v, res)
		}
	}
}
