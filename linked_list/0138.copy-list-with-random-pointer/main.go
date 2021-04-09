package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//Title: Copy List with Random Pointer
//Link: https://leetcode-cn.com/problems/copy-list-with-random-pointer

var m map[*Node]*Node

func copyRandomList(head *Node) *Node {
	m = map[*Node]*Node{}
	return dfs(head)
}

func dfs(node *Node) *Node {
	if node == nil {
		return nil
	}

	_, ok := m[node]
	if !ok {
		m[node] = &Node{Val: node.Val}
		m[node].Next = dfs(node.Next)
		m[node].Random = dfs(node.Random)
	}

	return m[node]
}
