package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var queue = []*TreeNode{root}
	for {
		counter := len(queue)
		var max = -1 << 63
		if counter == 0 {
			break
		}
		for _, v := range queue[:counter] {
			if max < v.Val {
				max = v.Val
			}
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		res = append(res, max)
		queue = queue[counter:]
	}

	return res
}
