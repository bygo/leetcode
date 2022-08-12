package main

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 有序链表转换二叉搜索树

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	numsL := len(nums)
	if numsL == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[numsL/2],
		Left:  sortedArrayToBST(nums[:numsL/2]), // 偶数时，左边多一个
		Right: sortedArrayToBST(nums[numsL/2+1:]),
	}
}
