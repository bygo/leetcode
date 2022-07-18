package main

// https://leetcode.cn/problems/n-ary-tree-level-order-traversal/

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

	var dep = -1
	var que = []*Node{root}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		dep++
		depsNums = append(depsNums, []int{})
		for _, node := range que[:queL] {
			depsNums[dep] = append(depsNums[dep], node.Val)
			for _, child := range node.Children {
				que = append(que, child)
			}
		}
		que = que[queL:]
	}
	return depsNums
}
