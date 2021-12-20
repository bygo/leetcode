package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack = []*Node{root}
	for 0 < len(stack) {
		res = append(res, root.Val)
		for _, n := range root.Children {
			stack = append(stack, n) //入栈
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	l := len(res) - 1
	for i := 0; i < l/2+1; i++ {
		res[i], res[l-i] = res[l-i], res[i]
	}
	return res
}
