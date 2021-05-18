package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res [][]int
	var level int
	for {
		l := len(queue)
		if 0 == l {
			break
		}

		res = append(res, []int{})

		for _, q := range queue {
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
			res[level] = append(res[level], q.Val)
		}
		queue = queue[l:]
		if level%2 == 1 {
			l, r := 0, len(res[level])-1
			for l < r {
				res[level][l], res[level][r] = res[level][r], res[level][l]
				l++
				r--
			}
		}
		level++
	}

	return res
}
