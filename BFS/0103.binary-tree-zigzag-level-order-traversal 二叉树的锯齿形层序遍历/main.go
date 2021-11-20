package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binary tree zigzag level order traversal
// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue = []*TreeNode{root}
	var depth, cnt int
	for 0 < len(queue) {
		cnt = len(queue)
		res = append(res, []int{})
		for _, q := range queue[:cnt] {
			res[depth] = append(res[depth], q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		if depth%2 == 1 {
			var i = 0
			var j = len(res[depth]) - 1
			for i < j {
				res[depth][i], res[depth][j] = res[depth][j], res[depth][i]
				i++
				j--
			}
		}
		depth++
		queue = queue[cnt:]
	}

	return res
}
