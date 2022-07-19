package main

type Node struct {
	Val  int
	Next *Node
}

// https://leetcode.cn/problems/insert-into-a-sorted-circular-linked-list

func insert(root *Node, x int) *Node {
	if root == nil {
		root := &Node{Val: x}
		root.Next = root
		return root
	}

	curr := root

	for {
		if curr.Val <= x && x <= curr.Next.Val || //升序中间 直接插入
			curr.Next.Val < curr.Val && (curr.Val <= x || x <= curr.Next.Val) || // curr < prev 时 必定 prev = tail & curr = head  ,如果正好可以插入 就插入
			curr.Next == root { // 如果都找不到了 所以值相同 随便插了
			curr.Next = &Node{Val: x, Next: curr.Next}
			break
		}
		curr = curr.Next
	}

	//for !(curr.Val <= x && x <= curr.Next.Val ||
	//	curr.Next.Val < curr.Val && (curr.Val <= x || x <= curr.Next.Val) ||
	//	curr.Next == root) {
	//	curr = curr.Next
	//}
	//curr.Next = &Node{Val: x, Next: curr.Next}

	return root
}
