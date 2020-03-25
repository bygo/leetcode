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
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(arr)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{nums[len(nums)/2], sortedArrayToBST(nums[:len(nums)/2]), sortedArrayToBST(nums[len(nums)/2+1:])}
}
