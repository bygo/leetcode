package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack = []*Node{root}
	for 0 < len(stack) {
		res = append(res, root.Val) //前序输出
		for i := len(root.Children) - 1; 0 <= i; i-- {
			stack = append(stack, root.Children[i]) //入栈
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}
