package main

// https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的垂直遍历

type Pair struct {
	node *TreeNode
	idx  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	idxMpNums := map[int][]int{}

	var que []*Pair
	que = append(que, &Pair{
		node: root,
		idx:  0,
	})
	for {
		queL := len(que)
		if queL == 0 {
			break
		}

		for _, pair := range que[:queL] {
			node := pair.node
			idx := pair.idx
			idxMpNums[idx] = append(idxMpNums[idx], node.Val)
			if node.Left != nil {
				que = append(que, &Pair{
					node: node.Left,
					idx:  idx - 1,
				})
			}
			if node.Right != nil {
				que = append(que, &Pair{
					node: node.Right,
					idx:  idx + 1,
				})
			}
		}
		que = que[queL:]
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
