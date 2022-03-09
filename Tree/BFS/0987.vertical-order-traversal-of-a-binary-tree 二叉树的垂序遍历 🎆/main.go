package main

import "sort"

// https://leetcode-cn.com/problems/vertical-order-traversal-of-a-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树的垂序遍历

type Pair struct {
	node *TreeNode
	row  int
	col  int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	que := []*Pair{}
	var dfs func(node *TreeNode, row, col int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		que = append(que, &Pair{
			node: node,
			row:  row,
			col:  col,
		})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)
	sort.Slice(que, func(i, j int) bool {
		a, b := que[i], que[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.node.Val < b.node.Val)
	})

	var depsNums [][]int
	var col = -1001
	var idx = -1
	for _, pair := range que {
		if pair.col != col {
			idx++
			col = pair.col
			depsNums = append(depsNums, []int{})
		}
		depsNums[idx] = append(depsNums[idx], pair.node.Val)
	}
	return depsNums
}
