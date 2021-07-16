package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode-cn.com/problems/copy-list-with-random-pointer

var m map[*Node]*Node

func copyRandomList2(head *Node) *Node {
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

// O(1) 空间
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 一遍遍历创建节点
	old := head
	for old != nil {
		n := &Node{Val: old.Val}
		n.Next = old.Next
		old.Next = n
		old = n.Next
	}

	// 二遍遍历 补全random
	old = head
	for old != nil {
		if old.Random != nil {
			old.Next.Random = old.Random.Next
		}
		old = old.Next.Next
	}

	// 三遍遍历 删除old
	old = head
	n := head.Next
	newHead := n
	for old != nil {
		old.Next = old.Next.Next
		if n.Next != nil {
			n.Next = n.Next.Next
		}
		old = old.Next
		n = n.Next
	}
	return newHead
}
