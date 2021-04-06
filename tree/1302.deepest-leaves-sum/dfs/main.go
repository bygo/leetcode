package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int
var sum int

func deepestLeavesSum(root *TreeNode) int {
	max, sum = 0, 0
	dfs(root, 0)
	return sum
}

func dfs(root *TreeNode, level int) {
	if root != nil {
		if max < level {
			max = level
			sum = root.Val
		} else if max == level {
			sum += root.Val
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}

func bfs(root *TreeNode) int {
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
		res = 0
		for _, q := range queue[:counter] {
			res += q.Val
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[counter:]
	}
	return res
}
