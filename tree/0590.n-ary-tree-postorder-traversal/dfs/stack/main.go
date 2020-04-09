package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res = []int{}
	var stack = []*Node{root}
	for 0 < len(stack) {
		root = stack[len(stack)-1]
		res = append(res, root.Val)
		stack = stack[:len(stack)-1]
		l := len(root.Children)
		for i := 0; i < l; i++ {
			stack = append(stack, root.Children[i]) //入栈
		}
	}

	l := len(res) - 1
	for i := 0; i < l/2; i++ {
		res[i], res[l-i] = res[l-i], res[i]
	}
	return res
}
