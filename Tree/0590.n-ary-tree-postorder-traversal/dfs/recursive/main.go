package main

type Node struct {
	Val      int
	Children []*Node
}

var res []int

func postorder(root *Node) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *Node) {
	if root != nil {
		for _, n := range root.Children {
			dfs(n)
		}
		res = append(res, root.Val) //后序输出
	}
}
