package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isValidBST(root *TreeNode) bool {
//	var last = &TreeNode{Val: -1 << 63}
//	var stack []*TreeNode
//	for root != nil || len(stack) > 0 {
//		//push
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//
//		//pop
//		pre := len(stack) - 1
//		if last.Val >= stack[pre].Val {
//			return false
//		}
//		last = stack[pre]
//		root = stack[pre].Right
//		stack = stack[:pre]
//	}
//	return true
//}

func isValid(root *TreeNode) bool {
	var last = &TreeNode{Val: -1 << 63}
	var stack []*TreeNode
	for 0 < len(stack) || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := len(stack) - 1
		if stack[top].Val < last.Val {
			return false
		}
		last = stack[top]
		root = stack[top].Right
		stack = stack[:top]
	}
	return true
}
