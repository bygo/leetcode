package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res = []int{}
	var stack = []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			res = append(res, root.Val)
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i]) //入栈
			}
			root = root.Children[0]
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	//反转 变成后序遍历 左右根
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
