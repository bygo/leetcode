package _0_11_12

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	var level int
	for {
		counter := len(queue)
		if counter == 0 {
			break
		}

		res = append(res, []int{})
		for _, v := range queue[:counter] {
			res[level] = append(res[level], v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}

		queue = queue[:counter]
		level += 1
	}
	return res
}
