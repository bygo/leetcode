package main

// https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	m := map[int][]int{}

	var queue = []*TreeNode{root}
	var index = []int{0}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := range queue[:cnt] {
			r := queue[i]
			k := index[i]
			m[index[i]] = append(m[index[i]], r.Val)
			if r.Left != nil {
				queue = append(queue, r.Left)
				index = append(index, k-1)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
				index = append(index, k+1)
			}
		}
		queue = queue[cnt:]
		index = index[cnt:]
	}

	var res [][]int
	for i := -100; i < 100; i++ {
		v, ok := m[i]
		if ok {
			res = append(res, v)
		}
	}
	return res
}
