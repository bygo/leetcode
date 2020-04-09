package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val) //前序输出
			l := len(root.Children)
			if 0 < l {
				for i := l - 1; 0 < i; i-- {
					stack = append(stack, root.Children[i]) //入栈
				}
				root = root.Children[0]
			} else {
				break
			}
		}
		root = stack[len(stack)-1] //出栈
		stack = stack[:len(stack)-1]
	}
	return res
}
