package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}
		res = queue[0].Val
		for _, v := range queue[:counter] {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[counter:]
	}
	return res
}
