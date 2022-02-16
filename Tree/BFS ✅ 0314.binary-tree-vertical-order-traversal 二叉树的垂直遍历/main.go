package main

// https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的垂直遍历

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	idxMpNums := map[int][]int{}

	var que = []*TreeNode{root}
	var idxes = []int{0}
	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}

		for i, q := range que[:cnt] {
			idx := idxes[i]
			idxMpNums[idx] = append(idxMpNums[idx], q.Val)
			if q.Left != nil {
				que = append(que, q.Left)
				idxes = append(idxes, idx-1)
			}
			if q.Right != nil {
				que = append(que, q.Right)
				idxes = append(idxes, idx+1)
			}
		}
		que = que[cnt:]
		idxes = idxes[cnt:]
	}

	var depsNums [][]int
	for idx := -100; idx < 100; idx++ {
		nums, ok := idxMpNums[idx]
		if ok {
			depsNums = append(depsNums, nums)
		}
	}
	return depsNums
}
