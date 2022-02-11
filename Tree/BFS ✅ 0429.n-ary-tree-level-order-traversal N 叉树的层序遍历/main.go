package main

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

type Node struct {
	Val      int
	Children []*Node
}

// ❓ N 叉树的层序遍历

func levelOrder(root *Node) [][]int {
	var depsNums [][]int
	if root == nil {
		return depsNums
	}

	var dep int
	var que = []*Node{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		depsNums = append(depsNums, []int{})
		for _, q := range que[:queL] {
			depsNums[dep] = append(depsNums[dep], q.Val)
			for _, child := range q.Children {
				que = append(que, child)
			}
		}
		dep++
		que = que[queL:]
	}
	return depsNums
}
