package main

type Node struct {
	Val  int
	Next *Node
}

//Link: https://leetcode-cn.com/problems/insert-into-a-sorted-circular-linked-list

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode := &Node{Val: x}
		aNode.Next = aNode
		return aNode
	}

	curr := aNode

	for {
		if curr.Val <= x && x <= curr.Next.Val || //升序中间 直接插入
			curr.Next.Val < curr.Val && (curr.Val <= x || x <= curr.Next.Val) || // curr < prev 时 必定 prev = tail & curr = head  ,如果正好可以插入 就插入
			curr.Next == aNode {
			curr.Next = &Node{Val: x, Next: curr.Next}
			break
		}
		curr = curr.Next
	}

	//for !(curr.Val <= x && x <= curr.Next.Val ||
	//	curr.Next.Val < curr.Val && (curr.Val <= x || x <= curr.Next.Val) ||
	//	curr.Next == aNode) {
	//	curr = curr.Next
	//}
	//curr.Next = &Node{Val: x, Next: curr.Next}

	return aNode
}
