package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := findMiddle(head) //find root
	node := &TreeNode{
		Val: mid.Val,
	}

	if head == mid { //just one element
		return node
	}

	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)
	return node
}

func findMiddle(node *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := node, node
	for fast != nil && fast.Next != nil { //end or end.Next
		prev = slow
		slow = slow.Next      //move 1
		fast = fast.Next.Next //move 2
	}

	if prev != nil {
		prev.Next = nil
	}
	return slow // = length/2 + 1
	//1=>1
	//2=>2
	//3=>2
	//4=>3
	//5=>3
	//6=>4
	//7=>4
}
