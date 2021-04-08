package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Random *Node
}

type NodeCopy struct {
	Val    int
	Left   *NodeCopy
	Right  *NodeCopy
	Random *NodeCopy
}

//Title: Clone Binary Tree With Random Pointer
//Link: https://leetcode-cn.com/problems/clone-binary-tree-with-random-pointer

var m = map[*Node]*NodeCopy{}

func copyRandomBinaryTree(root *Node) *NodeCopy {
	m = map[*Node]*NodeCopy{}
	return dfs(root)
}

func dfs(root *Node) *NodeCopy {
	if root == nil {
		return nil
	}
	_, ok := m[root]
	if !ok {
		m[root] = &NodeCopy{Val: root.Val}
		m[root].Left = dfs(root.Left)
		m[root].Right = dfs(root.Right)
		m[root].Random = dfs(root.Random)
	}

	return m[root]
}
