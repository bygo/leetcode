package main

type Node struct {
	Val      int
	Children []*Node
}

// n-ary-tree level order traversal
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var depth int
	var queue = []*Node{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, []int{})
		for _, v := range queue[:cnt] {
			res[depth] = append(res[depth], v.Val)
			for _, c := range v.Children {
				queue = append(queue, c)
			}
		}
		depth++
		queue = queue[cnt:]
	}
	return res
}
