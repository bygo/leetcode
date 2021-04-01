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
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[l/2],
		Left:  sortedArrayToBST(nums[:l/2]),
		Right: sortedArrayToBST(nums[l/2+1:]),
	}
}
